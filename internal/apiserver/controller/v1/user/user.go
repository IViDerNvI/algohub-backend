package user

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type UserController struct {
	Srv service.Service
}

func NewUserController(store store.Store, cache cache.Cache, s3 substore.SubStore) *UserController {
	return &UserController{Srv: service.NewService(store, cache, s3)}
}
