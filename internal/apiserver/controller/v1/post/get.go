package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *PostController) Get(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	post, err := c.Service.Posts().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, post)
}
