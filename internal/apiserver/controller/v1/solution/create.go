package solution

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/sirupsen/logrus"
)

func (c *SolutionController) Create(ctx *gin.Context) {
	var solution v1.Solution

	if err := ctx.ShouldBindJSON(&solution); err != nil {
		core.WriteResponse(ctx, err, nil)
		return
	}

	operatorNameRaw, ok := ctx.Get("X-Operation-User-Name")
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	operatorName, ok := operatorNameRaw.(string)
	if !ok {
		core.WriteResponse(ctx, core.ErrNoAuthorization, nil)
		return
	}

	solution.ProblemID = ctx.Param("id")
	solution.Provider = operatorName

	logrus.Warnf("Solution endpoint: %v", solution)

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
