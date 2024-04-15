package RandFuncs

import (
	"context"
	"sqout/libs/State"

	"github.com/gin-gonic/gin"
)

func GetContext(c context.Context) *State.State {
	return c.(*gin.Context).MustGet("s").(*State.State)
}
