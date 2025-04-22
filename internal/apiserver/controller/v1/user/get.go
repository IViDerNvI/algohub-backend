package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *UserController) Get(ctx *gin.Context) {
	username := ctx.Param("id")
	if username == "" {
		core.WriteResponse(ctx, core.ErrUserNameNeed, nil)
		return
	}

	user, err := c.Srv.Users().Get(ctx, username, nil)
	if err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseGet, nil)
		return
	}

	user.Password = ""

	core.WriteResponse(ctx, nil, user)
}
