package subscribe

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubscribeController) Get(ctx *gin.Context) {
	rsrcId := ctx.Param("resourceid")
	rsrcType := ctx.Param("type")

	sub := &v1.Subscribe{
		ItemType: rsrcType,
		ItemName: rsrcId,
	}

	subList, err := c.Service.Subscribes().Get(ctx, sub, &v1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, subList)
}
