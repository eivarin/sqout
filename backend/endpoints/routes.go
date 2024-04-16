package endpoints

import (
	"sqout/endpoints/modules"
	"sqout/endpoints/probes"
	"github.com/gin-gonic/gin"
	"sqout/libs/State"
)

func SetupRoutes(r *gin.Engine, s *State.State) {
	modules.SetupRoutes(r, s)
	probes.SetupRoutes(r, s)
}
