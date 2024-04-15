package endpoints

import (
	"fmt"
	"net/http"
	"net/url"
	"sqout/libs/ModuleConfig"

	"github.com/gin-gonic/gin"
)

func SetupModulesRoutes(r *gin.Engine) {
	r.GET("/modules", get)
	r.GET("/modules/:name", getOne)
	r.DELETE("/modules/:name", delete)
	r.POST("/modules", post)
	r.PUT("/modules", put)
}

func get(ctx *gin.Context) {
	list, _ := ModuleConfig.GetAllModules(ctx)
	ctx.JSON(http.StatusOK, list)
}

func getOne(ctx *gin.Context) {
	name := ctx.Param("name")
	// replace %2f with / to allow for nested paths
	name, _ = url.PathUnescape(name)
	fmt.Printf("Getting module: %s\n", name)
	module, err := ModuleConfig.GetOneModule(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}
	ctx.JSON(http.StatusOK, module)
}

type PutBody struct {
	Name   string `json:"Name"`
	Branch string `json:"Branch"`
	Commit string `json:"Commit"`
}

func put(ctx *gin.Context) {
	var body PutBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}
	err := ModuleConfig.Update(ctx, body.Name, body.Branch, body.Commit)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}
	ctx.JSON(http.StatusOK, "Module updated successfully!")
}


func delete(c *gin.Context) {
	name := c.Param("name")
	// replace %2f with / to allow for nested paths
	name, _ = url.PathUnescape(name)
	fmt.Printf("Deleting module: %s\n", name)
	err := ModuleConfig.Delete(c, name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}
	c.JSON(http.StatusOK, "Module deleted successfully!")
}

type PostBody struct {
	Name   string `json:"Name"`
	Branch string `json:"Branch"`
	Commit string `json:"Commit"`
}

func post(ctx *gin.Context) {
	var body PostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}

	ModuleConfig.AddNewModule(ctx, body.Name, body.Branch, body.Commit)

	ctx.JSON(http.StatusOK, "Module added successfully!")
}
