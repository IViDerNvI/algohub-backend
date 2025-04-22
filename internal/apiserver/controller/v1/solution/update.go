package solution

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SolutionController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	old, err := c.Service.Solutions().Get(ctx, uint(id), nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var solution v1.Solution
	if err := ctx.ShouldBindJSON(&solution); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if old.Override(&solution).Validate() != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if err := old.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Solutions().Update(ctx, old, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseUpdate, nil)
		return
	}
	core.WriteResponse(ctx, nil, nil)
}
