package solution

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SolutionController) Create(ctx *gin.Context) {
	var solution v1.Solution

	if err := ctx.ShouldBindJSON(&solution); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName, ok := ctx.Get("X-Operation-User-Name")
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	solution.ProblemID, ok = operatorName.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := solution.Validate(); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if err := c.Service.Solutions().Create(ctx, &solution, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, solution)
}
