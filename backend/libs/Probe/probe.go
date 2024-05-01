package Probe

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sqout/libs/DbApi"
	"sqout/libs/Grafana"
	"sqout/libs/ModuleConfig"
	"sqout/libs/TimersMap"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @Probe struct

type Probe struct {
	Name             string            `bson:"_id"`
	Description      string            `bson:"description"`
	Options          map[string]string `bson:"options"`
	HeartbitInterval int               `bson:"heartbitInterval"`
	Module           string            `bson:"module"`
	Results          []bson.M          `bson:"results"`
	Alive            bool              `bson:"alive"`
}

func NewProbe(ctx context.Context, mCol *DbApi.ColFacade, pCol *DbApi.ColFacade, Timers *TimersMap.TimersMap, gs *Grafana.GrafanaState, name string, description string, options map[string]string, heartbitInterval int, moduleName string) error {
	moduleconfig, err := ModuleConfig.GetOneModule(ctx, mCol, moduleName)
	if err != nil {
		return err
	}
	if checkOptions(options, moduleconfig) {
		return errors.New("options do not match module flags")
	}
	if moduleconfig.Exe.KeepAlive {
		heartbitInterval = -1
	}
	n := Probe{
		Name:             name,
		Description:      description,
		Options:          options,
		HeartbitInterval: heartbitInterval,
		Module:           moduleName,
		Results:          []bson.M{},
		Alive:            true,
	}
	_, err = pCol.Col.InsertOne(ctx, n)
	go ActivateProbe(ctx, mCol, pCol, Timers, name)
	gs.LoadDashboardForProbe(moduleconfig.Path+"/dashboard.json", n.Name)
	return err
}

func GetOneProbe(ctx context.Context, pCol *DbApi.ColFacade, name string, includeResults bool) (Probe, error) {
	var p Probe
	var err error
	if includeResults {
		err = pCol.Col.FindOne(ctx, bson.M{"_id": name}).Decode(&p)
	} else {
		opts := options.FindOne().SetProjection(bson.M{"results": 0})
		err = pCol.Col.FindOne(ctx, bson.M{"_id": name}, opts).Decode(&p)
	}
	return p, err
}

func GetAllProbes(ctx context.Context, pCol *DbApi.ColFacade, includeResults bool) ([]Probe, error) {
	var probes []Probe
	var err error
	if includeResults {
		cur, err := pCol.Col.Find(ctx, bson.M{})
		if err != nil {
			return probes, err
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var p Probe
			err := cur.Decode(&p)
			if err != nil {
				return probes, err
			}
			probes = append(probes, p)
		}
	} else {
		opts := options.Find().SetProjection(bson.M{"results": 0})
		cur, err := pCol.Col.Find(ctx, bson.M{}, opts)
		if err != nil {
			return probes, err
		}
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var p Probe
			err := cur.Decode(&p)
			if err != nil {
				return probes, err
			}
			probes = append(probes, p)
		}
	}
	return probes, err
}

func DeleteProbe(ctx context.Context, pCol *DbApi.ColFacade, gs *Grafana.GrafanaState, name string) error {
	gs.DeleteDashboardOnGrafana(name)
	_, err := pCol.Col.DeleteOne(ctx, bson.M{"_id": name})
	return err
}

func RestartAllProbes(ctx context.Context, mCol *DbApi.ColFacade, pCol *DbApi.ColFacade, Timers *TimersMap.TimersMap, gs *Grafana.GrafanaState) {
	probes, _ := GetAllProbes(ctx, pCol, false)
	fmt.Println("Restarting all probes")
	for _, p := range probes {
		fmt.Printf("Restarting probe %s\n", p.Name)
		if p.Alive {
			go ActivateProbe(ctx, mCol, pCol, Timers, p.Name)
		}
		mc, _ := ModuleConfig.GetOneModule(ctx, mCol, p.Module)
		gs.LoadDashboardForProbe(mc.Path+"/dashboard.json", p.Name)
	}
}

func ActivateProbe(ctx context.Context, mCol *DbApi.ColFacade, pCol *DbApi.ColFacade, ts *TimersMap.TimersMap, name string) error {
	p, err := GetOneProbe(ctx, pCol, name, false)
	for p.Alive {
		RunProbeCMD(ctx, mCol, pCol, p)
		ts.WaitFor(name, p.HeartbitInterval)
		p, err = GetOneProbe(ctx, pCol, name, false)
		if err != nil {
			return err
		}
	}
	return err
}

func RunProbeCMD(ctx context.Context, mCol *DbApi.ColFacade, pCol *DbApi.ColFacade, p Probe) error {
	fmt.Printf("Running probe %s\n", p.Name)
	var args []string
	mc, err := ModuleConfig.GetOneModule(ctx, mCol, p.Module)
	if err != nil {
		return err
	}
	for _, flagName := range mc.Exe.FlagsOrder {
		if flagValue, ok := p.Options[flagName]; ok {
			if mc.Exe.Flags[flagName].Prefix != "" {
				args = append(args, mc.Exe.Flags[flagName].Prefix)
			}
			if !mc.Exe.Flags[flagName].IsEmpty {
				args = append(args, flagValue)
			}
		}
	}
	c1 := exec.Command(mc.Exe.CommandName, args...)
	fmt.Printf("Running command: %s\n", c1.String())
	c2 := exec.Command("bash", "-c", mc.Path+"/parse.sh")
	fmt.Printf("Running command: %s\n", c2.String())
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	var b2, e2 bytes.Buffer
	c2.Stdout = &b2
	c2.Stderr = &e2
	_ = c1.Start()
	_ = c2.Start()
	c1.Wait()
	w.Close()
	_ = c2.Wait()
	// fmt.Printf("Command output: %s\n", b2.String())
	// fmt.Printf("Command error: %s\n", e2.String())
	var result bson.A
	_ = bson.UnmarshalExtJSON(b2.Bytes(), true, &result)
	err = AddResult(ctx, pCol, p.Name, result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func AddResult(ctx context.Context, pCol *DbApi.ColFacade, name string, result bson.A) error {
	timestamp := time.Now()
	_, err := pCol.Col.UpdateOne(ctx, bson.M{"_id": name}, bson.M{"$push": bson.M{"results": bson.M{"timestamp": timestamp, "result": result}}})
	return err
}

func checkOptions(options map[string]string, mc ModuleConfig.ModuleConfig) bool {
	failed := false
	var optionsCopy = make(map[string]string)
	for k, v := range options {
		optionsCopy[k] = v
	}
	for flagName, flag := range mc.Exe.Flags {
		if flag.Required && optionsCopy[flagName] == "" {
			failed = true
			fmt.Printf("Flag %s is required\n", flagName)
		}
		delete(optionsCopy, flagName)
	}
	if len(optionsCopy) > 0 {
		failed = true
		for flagName := range optionsCopy {
			fmt.Printf("Flag %s is not a valid flag\n", flagName)
		}
	}
	return failed
}
