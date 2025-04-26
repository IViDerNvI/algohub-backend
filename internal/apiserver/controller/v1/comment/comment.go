package comment

import (
	"github.com/ividernvi/algohub/internal/apiserver/cache"
	"github.com/ividernvi/algohub/internal/apiserver/service"
	"github.com/ividernvi/algohub/internal/apiserver/store"
	"github.com/ividernvi/algohub/internal/apiserver/substore"
)

type CommentController struct {
	Service service.Service
}

func NewCommentController(store store.Store, cache cache.Cache, s3 substore.SubStore) *CommentController {
	return &CommentController{Service: service.NewService(store, cache, s3)}
}
