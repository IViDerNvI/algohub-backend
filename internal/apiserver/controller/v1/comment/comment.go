package comment

import (
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
)

type CommentController struct {
	Service service.Service
}

func NewCommentController(store store.Store) *CommentController {
	return &CommentController{Service: service.NewService(store)}
}
