package endpoints

import (
	"fmt"
	"net/http"
	"net/url"
	"sqout/libs/ModuleConfig"

	"github.com/gin-gonic/gin"
)

func SetupModulesRoutes(r *gin.Engine) {
	r.GET("/modules", getModule)
	r.DELETE("/modules/:name", deleteModule)
	r.POST("/modules", postModule)
	r.PUT("/modules", putModule)
}

func getModule(ctx *gin.Context) {
	list, _ := ModuleConfig.GetAllModules(ctx)
	ctx.JSON(http.StatusOK, list)
}

type PutBody struct {
	Name   string `json:"Name"`
	Branch string `json:"Branch"`
	Commit string `json:"Commit"`
}

func putModule(ctx *gin.Context) {
	var body PutBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}
	ModuleConfig.Update(ctx, body.Name, body.Branch, body.Commit)
}


func deleteModule(c *gin.Context) {
	name := c.Param("name")
	// replace %2f with / to allow for nested paths
	name, _ = url.PathUnescape(name)
	fmt.Printf("Deleting module: %s\n", name)
	ModuleConfig.Delete(c, name)
}

type PostBody struct {
	Name   string `json:"Name"`
	Branch string `json:"Branch"`
	Commit string `json:"Commit"`
}

func postModule(ctx *gin.Context) {
	var body PostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}

	ModuleConfig.AddNewModule(ctx, body.Name, body.Branch, body.Commit)

	ctx.JSON(http.StatusOK, "Module added successfully!")
}
