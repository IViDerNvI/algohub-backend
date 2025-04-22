package like

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *LikeController) Delete(ctx *gin.Context) {
	rsrcType := ctx.Param("type")
	rsrcId, err := strconv.Atoi(ctx.Param("resourceid"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		core.WriteResponse(ctx, core.ErrLoginNeed, nil)
		return
	}

	opUser, err := c.Service.Users().Verify(ctx, token, &v1.VerifyOptions{})
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	like := v1.Like{
		ItemType: rsrcType,
		ItemID:   uint(rsrcId),
		UserName: opUser.UserName,
	}

	if err := c.Service.Likes().Delete(ctx, &like, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
