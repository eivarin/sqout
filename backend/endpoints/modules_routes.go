package endpoints

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v2"
)

type ModuleConfig struct {
	id           string
	isGithubLink string `bson:isGithubLink`
	GitInfo      struct {
		RepoLink string `yaml:"repo-link"`
		Branch   string `yaml:"branch"`
		Commit   string `yaml:"commit"`
		Dev      bool   `yaml:"dev"`
	} `yaml:"git-info"`
	Exe struct {
		CommandName string `yaml:"command-name"`
		Flags       []Flag `yaml:"flags"`
	} `yaml:"exe"`
}

func (mc *ModuleConfig) toBSON() bson.M {
	return bson.M{
		"_id":          mc.id,
		"isGithubLink": mc.isGithubLink,
		"git-info": bson.M{
			"repo-link": mc.GitInfo.RepoLink,
			"branch":    mc.GitInfo.Branch,
			"commit":    mc.GitInfo.Commit,
			"dev":       mc.GitInfo.Dev,
		},
		"exe": bson.M{
			"command-name": mc.Exe.CommandName,
			"flags":        mc.Exe.Flags,
		},
	}
}

type Flag struct {
	Name  string `yaml:"name"`
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

func SetupModulesRoutes(r *gin.Engine) {
	r.GET("/modules", getModule)
	//r.DELETE("/modules/", deleteModule)
	r.POST("/modules", postModule)
}

func getModule(c *gin.Context) {
	coll := c.MustGet("collection").(*mongo.Collection)
	bla := coll.FindOne(c, bson.D{})
	var mod bson.M
	bla.Decode(&mod)
	fmt.Print(mod)
	c.JSON(http.StatusOK, mod)
}

func deleteModule(c *gin.Context) {
	coll := c.MustGet("collection").(*mongo.Collection)
	coll.DeleteOne(c, nil)
}

type PostBody struct {
	Name      string `json:"Name"`
	isGitLink bool   `json: "IsLink"`
}

func postModule(c *gin.Context) {
	coll := c.MustGet("collection").(*mongo.Collection)
	mod := new(ModuleConfig)
	var body PostBody
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA")
	}

	file, err := os.Open(fmt.Sprintf("modules/%s/config.yaml", body.Name))
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&mod)
	if err != nil {
		fmt.Print("BBBBBBBBBBBBBB")
	}
	mod.id = body.Name
	mod.isGithubLink = mod.isGithubLink

	res, _ := coll.InsertOne(c, mod.toBSON())
	id := res.InsertedID
	fmt.Println("Inserted ID:", id)

	c.JSON(http.StatusOK, "Module added successfully!")
}
