package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/internal/apiserver/controller/v1/user"
	v1 "github.com/ividernvi/algohub/model/v1"
	"github.com/ividernvi/algohub/pkg/core"
	"github.com/sirupsen/logrus"
)

func Authorize(userCtrl *user.UserController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("X-Operation-Login-Status", false)
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.Next()
			return
		}

		opUser, err := userCtrl.Srv.Users().Verify(ctx, token, &v1.VerifyOptions{})
		if err != nil {
			core.WriteResponse(ctx, core.ErrUserVerify, nil)
			ctx.Abort()
			return
		}

		ctx.Set("X-Operation-Login-Status", true)
		ctx.Set("X-Operation-User-ID", opUser.ObjMeta.ID)
		ctx.Set("X-Operation-User-InstanceID", opUser.InstanceID)
		ctx.Set("X-Operation-User-Name", opUser.UserName)
		ctx.Set("X-Operation-User-Status", opUser.Status)

		logrus.WithFields(logrus.Fields{
			"username": opUser.UserName,
			"status":   opUser.Status,
			"time":     time.Now(),
			"realm":    "internal/apiserver/middleware/user/authorize",
		}).Info("User authorized")

		ctx.Next()
	}
}
