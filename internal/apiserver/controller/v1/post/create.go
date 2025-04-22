package post

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *PostController) Create(ctx *gin.Context) {
	var post v1.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName := ctx.MustGet("X-Operation-User-Name")
	operatorStatus := ctx.MustGet("X-Operation-User-Status")

	if operatorStatus != "admin" && operatorName != post.Author {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := post.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Posts().Create(ctx, &post, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, post)
}
