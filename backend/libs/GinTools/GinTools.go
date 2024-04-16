package GinTools

import (
	"context"
	"sqout/endpoints"
	"sqout/libs/State"

	"github.com/gin-gonic/gin"
)

func InitGin() {
	r := gin.Default()
	s := State.InitState(context.Background())
	endpoints.SetupRoutes(r, s)
	r.UseRawPath = true
	r.UnescapePathValues = false
	r.Run() // listen and serve on
}
