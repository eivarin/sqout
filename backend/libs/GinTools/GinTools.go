package GinTools

import (
	"sqout/endpoints"
	"sqout/libs/State"

	"github.com/gin-gonic/gin"
)

func initContext(c *gin.Context) {
	s := State.InitState(c)
	c.Set("s", s)
}

func InitGin() {
	r := gin.Default()
	r.Use(initContext)
	endpoints.SetupModulesRoutes(r)
	r.UseRawPath = true
	r.UnescapePathValues = false
	r.Run() // listen and serve on
}
