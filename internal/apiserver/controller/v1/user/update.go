package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *UserController) Update(ctx *gin.Context) {
	requestUsername := ctx.Param("id")

	var user v1.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	opUserName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserName == nil {
		core.WriteResponse(ctx, core.ErrUnknownError, nil)
		return
	}

	opUserStatus, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatus == nil {
		core.WriteResponse(ctx, core.ErrUnknownError, nil)
		return
	}

	if opUserName != requestUsername && opUserStatus != "admin" {
		core.WriteResponse(ctx, core.ErrUnknownOperator, nil)
		return
	}

	user.UserName = requestUsername
	if c.Srv.Users().Update(ctx, &user, nil) != nil {
		core.WriteResponse(ctx, core.ErrUnknownError, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
