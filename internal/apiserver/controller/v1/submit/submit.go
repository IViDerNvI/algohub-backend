package submit

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type SubmitController struct {
	Service service.Service
}

func NewSubmitController(store store.Store, cache cache.Cache, s3 substore.SubStore) *SubmitController {
	return &SubmitController{
		Service: service.NewService(store, cache, s3),
	}
}
