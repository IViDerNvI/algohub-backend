package problem

import (
	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *ProblemController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	problem, err := c.Service.Problems().Get(ctx, id, nil)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, problem)
}
