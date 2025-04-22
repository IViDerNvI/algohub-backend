package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/ividernvi/algohub/pkg/util/jwtutil"
)

func (c *UserController) Login(ctx *gin.Context) {
	opUserNameRaw, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserNameRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserName, ok := opUserNameRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserStatusRaw, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatusRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserStatus, ok := opUserStatusRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserInstanceIdRaw, ok := ctx.Get("X-Operation-User-InstanceID")
	if !ok || opUserInstanceIdRaw == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}
	opUserInstanceId, ok := opUserInstanceIdRaw.(uint)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	jwt, err := jwtutil.CreateJWT(&v1.User{
		UserName: opUserName,
		Status:   opUserStatus,
		ObjMeta: v1.ObjMeta{
			InstanceID: opUserInstanceId,
		},
	})
	if err != nil {
		core.WriteResponse(ctx, core.ErrTokenCreateFailed, nil)
		return
	}

	ctx.Header("Authorization", jwt)
	core.WriteResponse(ctx, nil, jwt)
}
