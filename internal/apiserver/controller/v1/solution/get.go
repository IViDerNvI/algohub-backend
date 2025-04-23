package solution

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SolutionController) Get(ctx *gin.Context) {

	mapper := map[string]string{
		"problem_id": ctx.Query("problem_id"),
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Limit:    2,
		Offset:   0,
		Selector: selector,
	}
	listOptions.Complete()

	solution, err := c.Service.Solutions().List(ctx, listOptions)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, solution)
}
