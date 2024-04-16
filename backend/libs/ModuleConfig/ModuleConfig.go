package ModuleConfig

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sqout/libs/DbApi"

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
	KeepAlive   bool   `yaml:"keep-alive" bson:"keepAlive"`
	FlagsOrder []string `yaml:"flags-order" bson:"flagsOrder"`
	Flags       map[string]Flag `yaml:"flags" bson:"flags"`
}


type Flag struct {
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Required    bool   `yaml:"required"`
	Prefix      string `yaml:"prefix"`
}

func GetAllModules(ctx context.Context, mCol *DbApi.ColFacade) ([]ModuleConfig, error) {
	cursor, err := mCol.Col.Find(ctx, bson.D{})
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

func GetOneModule(ctx context.Context, mCol *DbApi.ColFacade, id string) (ModuleConfig, error) {
	var mc ModuleConfig
	err := mCol.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&mc)
	if err != nil {
		fmt.Printf("Error getting module with id %s\n", id)
		return mc, err
	}
	return mc, nil
}

func AddNewModule(ctx context.Context, mCol *DbApi.ColFacade, path string, branch string, commit string) error {
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

	res, _ := mCol.Col.InsertOne(ctx, mc)
	id := res.InsertedID
	fmt.Printf("Added module with id %v\n", id)
	return nil
}

func Update(ctx context.Context, mCol *DbApi.ColFacade, Name string, Branch string, Commit string) error {
	toChange, _ := GetOneModule(ctx, mCol, Name)
	err := toChange.ChangeVersion(Branch, Commit)
	if err != nil {
		return err
	}
	res := mCol.Col.FindOneAndReplace(ctx, bson.M{"_id": Name}, toChange)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func Delete(ctx context.Context, mCol *DbApi.ColFacade, name string) error {
	// var mc ModuleConfig
	res := mCol.Col.FindOneAndDelete(ctx, bson.M{"_id": name})
	if res.Err() != nil {
		return res.Err()
	}
	cmd := exec.Command("rm", "-rf", "./modules/"+name)
	cmd.Run()
	return nil
}

func runCMD(runningPath string, args []string) {
	command := exec.Command(args[0], args[1:]...)
	fmt.Printf("Running command: %s\n", command.String())
	command.Dir = runningPath
	command.Run()
}

func (mc *ModuleConfig) ChangeVersion(branch string, commit string) error {
	if mc.IsRepo {
		runningPath := "./modules/" + mc.Id
		if branch != "" || commit != "" {
			runCMD(runningPath, []string{"git", "checkout", "."})
			runCMD(runningPath, []string{"git", "fetch", "origin"})
		}
		if branch != ""{
			mc.GitInfo.Branch = branch
			runCMD(runningPath, []string{"git", "checkout", mc.GitInfo.Branch})
		}
		if commit != "" {
			mc.GitInfo.Commit = commit
			runCMD(runningPath, []string{"git", "checkout", mc.GitInfo.Commit})
		}
	}
	err := mc.Reload()
	if err != nil {
		return err
	}
	fmt.Printf("Changed module to branch %s and commit %s\n", branch, commit)
	return err
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
		fmt.Printf("Error: %v\n", err)
		return err
	}
	mc.Exe = parsed
	return nil
}
