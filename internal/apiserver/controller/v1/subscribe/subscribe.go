package subscribe

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type SubscribeController struct {
	Service service.Service
}

func NewSubscribeController(store store.Store, cache cache.Cache, s3 substore.SubStore) *SubscribeController {
	return &SubscribeController{Service: service.NewService(store, cache, s3)}
}
