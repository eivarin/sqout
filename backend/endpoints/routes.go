package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "sqout/docs"
	"sqout/endpoints/modules"
	"sqout/endpoints/probes"
	"sqout/libs/State"
)

func SetupRoutes(r *gin.Engine, s *State.State) {
	modules.SetupRoutes(r, s)
	probes.SetupRoutes(r, s)
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
