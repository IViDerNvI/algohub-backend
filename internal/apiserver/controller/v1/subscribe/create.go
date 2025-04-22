package subscribe

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubscribeController) Create(ctx *gin.Context) {
	rsrcId, err := strconv.Atoi(ctx.Param("resourceid"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	rsrcType := ctx.Param("type")

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

	sub := v1.Subscribe{
		ItemType: rsrcType,
		ItemID:   uint(rsrcId),
		UserName: opUserName,
	}

	if err := sub.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Subscribes().Create(ctx, &sub, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, sub)
}
