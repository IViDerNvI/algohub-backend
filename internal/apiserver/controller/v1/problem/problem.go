package problem

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type ProblemController struct {
	Service service.Service
}

func NewProblemController(store store.Store, cache cache.Cache, s3 substore.SubStore) *ProblemController {
	return &ProblemController{Service: service.NewService(store, cache, s3)}
}
