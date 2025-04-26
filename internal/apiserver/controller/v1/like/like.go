package like

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type LikeController struct {
	Service service.Service
}

func NewLikeController(store store.Store, cache cache.Cache, s3 substore.SubStore) *LikeController {
	return &LikeController{Service: service.NewService(store, cache, s3)}
}
