package like

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"

	v1 "github.com/ividernvi/algohub/model/v1"
)

func (c *LikeController) Get(ctx *gin.Context) {
	rsrcId, err := strconv.Atoi(ctx.Param("resourceid"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}
	rsrcType := ctx.Param("type")

	like := &v1.Like{
		ItemType: rsrcType,
		ItemID:   uint(rsrcId),
	}

	result, err := c.Service.Likes().Get(ctx, like, &v1.GetOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, result)
}
