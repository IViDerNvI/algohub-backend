package problem

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *ProblemController) Create(ctx *gin.Context) {
	var problem v1.Problem

	if err := ctx.ShouldBindJSON(&problem); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName, ok := ctx.Get("X-Operation-User-Name")
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	problem.Author, ok = operatorName.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := problem.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Problems().Create(ctx, &problem, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, problem)
}
