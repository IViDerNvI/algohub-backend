package subscribe

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubscribeController) Delete(ctx *gin.Context) {
	rsrcType := ctx.Param("type")
	rsrcId := ctx.Param("resourceid")

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

	sub := &v1.Subscribe{
		ItemType: rsrcType,
		ItemName: rsrcId,
		UserName: opUserName,
	}

	if err := c.Service.Subscribes().Delete(ctx, sub, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
