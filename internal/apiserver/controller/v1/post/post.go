package post

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type PostController struct {
	Service service.Service
}

func NewPostController(store store.Store, cache cache.Cache, s3 substore.SubStore) *PostController {
	return &PostController{Service: service.NewService(store, cache, s3)}
}
