package solution

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SolutionController) List(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		core.WriteResponse(ctx, core.ErrInvalidParams, nil)
		return
	}

	mapper := map[string]string{
		"problem_id": ctx.Param("id"),
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Limit:    limit,
		Offset:   offset,
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
