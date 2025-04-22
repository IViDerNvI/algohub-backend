package problem

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *ProblemController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	old, err := c.Service.Problems().Get(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	var problem v1.Problem
	if err := ctx.ShouldBindJSON(&problem); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if old.Override(&problem).Validate() != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if err := old.Validate(); err != nil {
		core.WriteResponse(ctx, core.ErrJSONFormation, nil)
		return
	}

	if err := c.Service.Problems().Update(ctx, old, nil); err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseUpdate, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
