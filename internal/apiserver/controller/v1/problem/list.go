package problem

import (
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *ProblemController) List(ctx *gin.Context) {
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
		"unique_id": ctx.Query("unique_id"),
		"title":     ctx.Query("title"),
		"level":     ctx.Query("level"),
		"tag":       ctx.Query("tag"),
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Limit:    limit,
		Offset:   offset,
		Selector: selector,
	}
	listOptions.Complete()

	problem, err := c.Service.Problems().List(ctx, listOptions)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	core.WriteResponse(ctx, nil, problem)
}
