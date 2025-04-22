package problem

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *ProblemController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	problem, err := c.Service.Problems().Get(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorName, ok := ctx.Get("X-Operation-User-Name")
	if !ok || operatorName == nil {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	if err := c.Service.Problems().Delete(ctx, problem.Unique_ID, nil); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, nil)
}
