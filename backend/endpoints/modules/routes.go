package modules

import (
	"fmt"
	"net/http"
	"net/url"
	"sqout/libs/DbApi"
	"sqout/libs/ModuleConfig"
	"sqout/libs/State"

	"github.com/gin-gonic/gin"
)

type moduleState struct {
	ModulesCol *DbApi.ColFacade
}

func SetupRoutes(r *gin.Engine, s *State.State) {
	var ms moduleState
	ms.ModulesCol = &s.ModulesCol
	r.GET("/modules", ms.get)
	r.GET("/modules/:name", ms.getOne)
	r.DELETE("/modules/:name", ms.delete)
	r.POST("/modules", ms.post)
	r.PUT("/modules", ms.put)
}

// @Summary GET all the modules in the database
// @Description get JSON of all the modules
// @Produce  json
// @Tags modules
// @Success 200 {example} json  "array of modules"
// @Router /modules [get]
func (ms *moduleState) get(ctx *gin.Context) {
	list, _ := ModuleConfig.GetAllModules(ctx, ms.ModulesCol)
	ctx.JSON(http.StatusOK, list)
}

// @Summary GET specific module in the database
// @Description get JSON of the module
// @Produce  json
// @Tags modules
// @Params name string
// @Success 200 {example} json  "module"
// @Router /modules/{module_name} [get]
func (ms *moduleState) getOne(ctx *gin.Context) {
	name := ctx.Param("name")
	// replace %2f with / to allow for nested paths
	name, _ = url.PathUnescape(name)
	fmt.Printf("Getting module: %s\n", name)
	module, err := ModuleConfig.GetOneModule(ctx, ms.ModulesCol, name)
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

func (ms *moduleState) put(ctx *gin.Context) {
	var body PutBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}
	err := ModuleConfig.Update(ctx, ms.ModulesCol, body.Name, body.Branch, body.Commit)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Module not found"})
		return
	}
	ctx.JSON(http.StatusOK, "Module updated successfully!")
}

func (ms *moduleState) delete(c *gin.Context) {
	name := c.Param("name")
	// replace %2f with / to allow for nested paths
	name, _ = url.PathUnescape(name)
	fmt.Printf("Deleting module: %s\n", name)
	err := ModuleConfig.Delete(c, ms.ModulesCol, name)
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

func (ms *moduleState) post(ctx *gin.Context) {
	var body PostBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Print("AAAAAAAAAAAAAAAAA	")
	}

	ModuleConfig.AddNewModule(ctx, ms.ModulesCol, body.Name, body.Branch, body.Commit)

	ctx.JSON(http.StatusOK, "Module added successfully!")
}
