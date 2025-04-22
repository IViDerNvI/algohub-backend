package solution

import (
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type SolutionController struct {
	Service service.Service
}

func NewSolutionController(store store.Store) *SolutionController {
	return &SolutionController{
		Service: service.NewService(store),
	}
}
