package like

import (
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type LikeController struct {
	Service service.Service
}

func NewLikeController(store store.Store) *LikeController {
	return &LikeController{Service: service.NewService(store)}
}
