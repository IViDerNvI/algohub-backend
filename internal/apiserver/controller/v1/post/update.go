package post

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *PostController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	old, err := c.Service.Posts().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var post v1.Post
	if err := ctx.ShouldBindJSON(&post); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
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

	if opUserStatus != "admin" && opUserName != old.Author {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	old.Override(&post)
	if err := c.Service.Posts().Update(ctx, old, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseUpdate, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
