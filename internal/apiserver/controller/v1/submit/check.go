package submit

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *SubmitController) Check(ctx *gin.Context) {

	opUserNameRaw, ok := ctx.Get("X-Operation-User-Name")
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	opUserName, ok := opUserNameRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	mapper := map[string]string{
		"problem_id": ctx.Param("id"),
		"author":     opUserName,
		"status":     "ACCEPTED",
	}

	selector := v1.Selector(mapper)

	listOptions := &v1.ListOptions{
		Offset:   0,
		Limit:    10000000,
		Selector: selector,
	}
	listOptions.Complete()

	posts, err := c.Service.Submits().List(ctx, listOptions)
	if err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	if posts.TotalItems != 0 {
		core.WriteResponse(ctx, nil, true)
		return
	}

	core.WriteResponse(ctx, nil, false)
}
