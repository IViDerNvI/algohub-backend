package like

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
)

func (c *LikeController) Check(ctx *gin.Context) {
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
		"item_type": ctx.Param("type"),
		"username":  opUserName,
		"item_id":   ctx.Param("resourceid"),
	}

	selector := v1.Selector(mapper)

	result, err := c.Service.Likes().List(ctx, &v1.ListOptions{
		Offset:   0,
		Limit:    1,
		Selector: selector,
	})
	if err != nil {
		core.WriteResponse(ctx, core.ErrDatabaseQuery, nil)
		return
	}

	if result.TotalItems == 0 {
		core.WriteResponse(ctx, nil, false)
		return
	}

	core.WriteResponse(ctx, nil, true)
}
