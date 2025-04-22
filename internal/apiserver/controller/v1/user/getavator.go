package user

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (u *UserController) GetAvatar(ctx *gin.Context) {
	avatar, err := u.Srv.Users().GetAvatar(ctx, ctx.Param("id"), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if avatar == nil {
		core.WriteResponse(ctx, errors.New(""), nil)
		return
	}

	core.WriteResponseWithFile(ctx, nil, avatar)
}
