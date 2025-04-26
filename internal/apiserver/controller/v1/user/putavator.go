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

	if len(avatarBytes) > 10*1024*1024 {
		core.WriteResponse(ctx, core.ErrFileTooLarge, nil)
		return
	}

	if ctx.ContentType() != "image/png" {
		core.WriteResponse(ctx, core.ErrInvalidFileType, nil)
		return
	}

	opUserNameRaw, exists := ctx.Get("X-Operation-User-Name")
	if !exists {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserName, ok := opUserNameRaw.(string)
	if !ok {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if opUserName != ctx.Param("id") {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	avatarUrl, err := u.Srv.Users().PutAvatar(ctx, &v1.PutOptions{
		Realm:      "algohub-user",
		File_UUID:  ctx.Param("id"),
		File_Bytes: avatarBytes,
		File_type:  "image/png",
	})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if u.Srv.Users().Update(ctx, &v1.User{
		UserName: ctx.Param("id"),
		Avatar:   avatarUrl.SubUrl,
	}, &v1.UpdateOptions{}) != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, avatarUrl)
}
