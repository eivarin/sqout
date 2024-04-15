package ModuleConfig

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sqout/libs/RandFuncs"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/yaml.v2"
)

type ModuleConfig struct {
	Id      string `bson:"_id"`
	Path    string `bson:"path"`
	IsRepo  bool  `bson:"isRepo"`
	GitInfo struct {
		Branch string `bson:"branch"`
		Commit string `bson:"commit"`
	} `bson:"gitInfo"`
	Exe exe `bson:"exe"`
}

type exe struct {
	CommandName string `yaml:"command-name" bson:"commandName"`
	Description string `yaml:"description" bson:"description"`
	keepAlive   bool   `yaml:"keep-alive" bson:"keepAlive"`
	Flags       []Flag `yaml:"flags" bson:"flags"`
}


type Flag struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
	Prefix      string `yaml:"prefix"`
}

func GetAllModules(ctx context.Context) ([]ModuleConfig, error) {
	s := RandFuncs.GetContext(ctx)
	cursor, err := s.ModulesCol.Col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var modules []ModuleConfig
	if err = cursor.All(ctx, &modules); err != nil {
		fmt.Println("Error getting all modules")
		return nil, err
	}
	fmt.Printf("Got all modules: %v\n", modules)
	return modules, nil
}

func GetOneModule(ctx context.Context, id string) (ModuleConfig, error) {
	s := RandFuncs.GetContext(ctx)
	var mc ModuleConfig
	err := s.ModulesCol.Col.FindOne(ctx, bson.E{Key: "_id",Value:  id}).Decode(&mc)
	if err != nil {
		fmt.Printf("Error getting module with id %s\n", id)
		return mc, err
	}
	fmt.Printf("Got module with id %s: %v\n", id, mc)
	return mc, nil
}

func AddNewModule(ctx context.Context, path string, branch string, commit string) error {
	fmt.Println("Adding new module")
	regex := regexp.MustCompile(`(https://)?(www\.)?(github\.com/[a-zA-Z0-9-]+/[a-zA-Z0-9-]+)`)
	match := regex.FindStringSubmatch(path)
	var mc ModuleConfig
	mc.IsRepo = match != nil
	if !mc.IsRepo {
		mc.Id = path
	} else {
		mc.Id = match[3]
	}
	mc.Path = "./modules/" + mc.Id
	if _, err := os.Stat(mc.Path); os.IsNotExist(err) {
		os.MkdirAll(mc.Path, os.ModePerm)
		cmd := exec.Command("git", "clone", path, mc.Path)
		cmd.Run()
	}
	mc.ChangeVersion(branch, commit)

	//add to database
	s := RandFuncs.GetContext(ctx)
	res, _ := s.ModulesCol.Col.InsertOne(ctx, mc)
	id := res.InsertedID
	fmt.Printf("Added module with id %v\n", id)
	return nil
}

func Update(ctx context.Context, Name string, Branch string, Commit string) {
	s := RandFuncs.GetContext(ctx)
	toChange, _ := GetOneModule(ctx, Name)
	toChange.ChangeVersion(Branch, Commit)
	s.ModulesCol.Col.FindOneAndReplace(ctx, bson.E{Key: "_id", Value: Name}, toChange)
}

func Delete(ctx context.Context, name string) {
	s := RandFuncs.GetContext(ctx)
	s.ModulesCol.Col.FindOneAndDelete(ctx, bson.E{Key: "_id", Value: name})
	cmd := exec.Command("rm", "-rf", "./modules/"+name)
	cmd.Run()
}

func (mc *ModuleConfig) ChangeVersion(branch string, commit string) {
	if !mc.IsRepo {
		runningPath := "./modules/" + mc.Id
		cmd := exec.Command("git", "fetch", "origin")
		cmd.Dir = runningPath
		cmd.Run()
		if branch != ""{
			mc.GitInfo.Branch = branch
			cmd = exec.Command("git", "switch", mc.GitInfo.Branch)
			cmd.Dir = runningPath
			cmd.Run()
		}
		cmd = exec.Command("git", "pull")
		cmd.Dir = runningPath
		cmd.Run()
		if commit != "" {
			mc.GitInfo.Commit = commit
			cmd = exec.Command("git", "checkout", mc.GitInfo.Commit)
			cmd.Dir = runningPath
			cmd.Run()
		}
	}
	mc.Reload()
	fmt.Printf("Changed module to branch %s and commit %s\n", branch, commit)
}

func (mc *ModuleConfig) Reload() error {
	//check if necessary files exist
	filenames := []string{"config.yaml", "parse.sh", "install.sh"}
	fileExistence := make(map[string]bool)

	for _, filename := range filenames {
		_, err := os.Stat(mc.Path + "/" + filename)
		if err != nil {
			if os.IsNotExist(err) {
				fileExistence[filename] = false
			}
		} else {
			fileExistence[filename] = true
		}
	}

	allExists := true
	for filename, exists := range fileExistence {
		if !exists {
			fmt.Printf("File %s does not exist\n", filename)
			allExists = false
		}
	}
	if !allExists {
		return fmt.Errorf("not all necessary files exist")
	}
	//run install.sh
	cmdInstall := exec.Command("bash", mc.Path+"/install.sh")
	err := cmdInstall.Run()
	if err != nil {
		fmt.Println("Error running install.sh")
		return err
	}

	file, err := os.Open(mc.Path + "/config.yaml")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()

	var parsed exe
	err = yaml.NewDecoder(file).Decode(&parsed)
	if err != nil {
		fmt.Println("Error decoding yaml")
		return err
	}
	mc.Exe = parsed
	return nil
}
