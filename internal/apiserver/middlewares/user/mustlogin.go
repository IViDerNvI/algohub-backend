package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func MustLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginStatus, ok := ctx.Get("X-Operation-Login-Status")
		if !ok || loginStatus == false {
			core.WriteResponse(ctx, core.ErrLoginNeed, nil)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
