package submit

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubmitController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	old, err := c.Service.Submits().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var submit v1.Submit
	if err := ctx.ShouldBindJSON(&submit); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	opUserName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || opUserName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if opUserName != old.Author {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	old.Override(&submit)
	if err := c.Service.Submits().Update(ctx, old, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseUpdate, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
