package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *UserController) Delete(ctx *gin.Context) {
	username := ctx.Param("id")

	opUserName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserStatus, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatus == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if opUserStatus != "admin" && opUserName != username {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	err := c.Srv.Users().Delete(ctx, username, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
