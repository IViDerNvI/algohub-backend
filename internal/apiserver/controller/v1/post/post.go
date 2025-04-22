package post

import (
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type PostController struct {
	Service service.Service
}

func NewPostController(store store.Store) *PostController {
	return &PostController{Service: service.NewService(store)}
}
