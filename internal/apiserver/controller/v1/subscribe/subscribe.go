package subscribe

import (
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type SubscribeController struct {
	Service service.Service
}

func NewSubscribeController(store store.Store) *SubscribeController {
	return &SubscribeController{Service: service.NewService(store)}
}
