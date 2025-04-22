package subscribe

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubscribeController) Get(ctx *gin.Context) {
	rsrcId, err := strconv.Atoi(ctx.Param("resourceid"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}
	rsrcType := ctx.Param("type")

	sub := &v1.Subscribe{
		ItemType: rsrcType,
		ItemID:   uint(rsrcId),
	}

	result, err := c.Service.Subscribes().Get(ctx, sub, &v1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, result)
}
