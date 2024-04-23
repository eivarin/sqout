package probes

import (
	"net/http"
	"sqout/libs/DbApi"
	"sqout/libs/Probe"
	"sqout/libs/State"
	"sqout/libs/TimersMap"

	"github.com/gin-gonic/gin"
)

// Custom response model for Swagger documentation because of bson.A
type ProbeResponse struct {
	Name             string            `json:"_id"`
	Description      string            `json:"description"`
	Options          map[string]string `json:"options"`
	HeartbitInterval int               `json:"heartbitInterval"`
	Module           string            `json:"module"`
	Results          []interface{}     `json:"results"` // swagger can't find type definitiion bson.a
	Alive            bool              `json:"alive"`
}

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
// @Description Create a new probe and run it
// @Tags probes
// @Accept json
// @Produce json
// @Param body body postBody true "Probe information"
// @Success 200 {string} string "Probe created successfully!"
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

// @Summary GET all probes
// @Description Retrieve all probes
// @Tags probes
// @Accept json
// @Produce json
// @Param includeResults query bool false "Include results in response"
// @Success 200 {object} []ProbeResponse "List of probes"
// @Router /probes [get]
func (ps *probeState) get(ctx *gin.Context) {
	includeResults := ctx.Query("includeResults") == "true"
	probes, err := Probe.GetAllProbes(ctx, ps.ProbesCol, includeResults)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if probes == nil {
		probes = []Probe.Probe{}
	}
	ctx.JSON(http.StatusOK, probes)
}

// @Summary GET a specific probe by name
// @Description Retrieve a specific probe by its name
// @Tags probes
// @Accept json
// @Produce json
// @Param name path string true "Probe Name"
// @Param includeResults query bool false "Include results in response"
// @Success 200 {object} ProbeResponse "Probe details"
// @Router /probes/{name} [get]
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

// @Summary DELETE a probe by name
// @Description Delete a probe by its name
// @Tags probes
// @Accept json
// @Produce json
// @Param name path string true "Probe Name"
// @Success 200 {string} string "Probe deleted successfully!"
// @Router /probes/{name} [delete]
func (ps *probeState) delete(c *gin.Context) {
	name := c.Param("name")
	err := Probe.DeleteProbe(c, ps.ProbesCol, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Probe deleted successfully!")
}
