package probes

import (
	"net/http"
	"sqout/libs/DbApi"
	"sqout/libs/Probe"
	"sqout/libs/State"
	"sqout/libs/TimersMap"

	"github.com/gin-gonic/gin"
)

type probeState struct {
	ModulesCol *DbApi.ColFacade
	ProbesCol  *DbApi.ColFacade
	Timers     *TimersMap.TimersMap
}

func SetupRoutes(r *gin.Engine, s *State.State) {
	var ps probeState
	ps.ModulesCol = &s.ModulesCol
	ps.ProbesCol = &s.ProbesCol
	ps.Timers = &s.Timers
	r.POST("/probes", ps.post)
	r.GET("/probes", ps.get)
	r.GET("/probes/:name", ps.getOne)
	r.DELETE("/probes/:name", ps.delete)
}

type postBody struct {
	Name             string            `bson:"_id"`
	Description      string            `bson:"description"`
	Options          map[string]string `bson:"options"`
	HeartbitInterval int               `bson:"heartbitInterval"`
	ModuleName       string            `bson:"module"`
}

// @Summary POST a new probe on the database
// @Description Run the probe
// @Tags probes
// @Accept json
// @Prouce json
// @Success 200 {example} json  "probe"
// @Router /probes [post]
func (ps *probeState) post(ctx *gin.Context) {
	var body postBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := Probe.NewProbe(ctx, ps.ModulesCol, ps.ProbesCol, ps.Timers, body.Name, body.Description, body.Options, body.HeartbitInterval, body.ModuleName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Probe created successfully!")
}

func (ps *probeState) get(ctx *gin.Context) {
	includeResults := ctx.Query("includeResults") == "true"
	probes, err := Probe.GetAllProbes(ctx, ps.ProbesCol, includeResults)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, probes)
}

func (ps *probeState) getOne(ctx *gin.Context) {
	name := ctx.Param("name")
	includeResults := ctx.Query("includeResults") == "true"
	probe, err := Probe.GetOneProbe(ctx, ps.ProbesCol, name, includeResults)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, probe)
}

func (ps *probeState) delete(c *gin.Context) {
	name := c.Param("name")
	err := Probe.DeleteProbe(c, ps.ProbesCol, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Probe deleted successfully!")
}
