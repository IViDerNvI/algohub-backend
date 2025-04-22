package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func MustAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		opUserStatus, ok := ctx.Get("X-Operation-User-Status")
		if !ok {
			core.WriteResponse(ctx, core.ErrUnknownOperator, nil)
			ctx.Abort()
		}
		if opUserStatus != "admin" {
			core.WriteResponse(ctx, core.ErrAdminNeed, nil)
			ctx.Abort()
		}
		ctx.Next()
	}
}
