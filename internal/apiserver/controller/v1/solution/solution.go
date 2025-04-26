package solution

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type SolutionController struct {
	Service service.Service
}

func NewSolutionController(store store.Store, cache cache.Cache, s3 substore.SubStore) *SolutionController {
	return &SolutionController{
		Service: service.NewService(store, cache, s3),
	}
}
