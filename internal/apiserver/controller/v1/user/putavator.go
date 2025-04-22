package user

import (
	"io"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (u *UserController) PutAvatar(ctx *gin.Context) {
	avatarBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if len(avatarBytes) > 5*1024*1024 {
		core.WriteResponse(ctx, core.ErrFileTooLarge, nil)
		return
	}

	if ctx.ContentType() != "image/png" {
		core.WriteResponse(ctx, core.ErrInvalidFileType, nil)
		return
	}

	err = u.Srv.Users().PutAvatar(ctx, ctx.Param("id"), avatarBytes, &v1.UpdateOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
