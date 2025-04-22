package comment

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *CommentController) Create(ctx *gin.Context) {
	var comment v1.Comment

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	opUserName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserStatus, ok := ctx.Get("X-Operation-User-Status")
	if !ok || opUserStatus == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if opUserStatus != "admin" && opUserName != comment.Auhtor {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := comment.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Comments().Create(ctx, &comment, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	ctx.JSON(200, comment)
}
