package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *PostController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	post, err := c.Service.Posts().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || operatorName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	operatorStatus, ok := ctx.Get("X-Operation-User-Status")
	if !ok || operatorStatus == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if operatorStatus != "admin" && operatorName != post.Author {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := c.Service.Posts().Delete(ctx, uint(id), nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
