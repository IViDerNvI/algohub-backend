package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type CommentService interface {
	Create(ctx context.Context, comment *v1.Comment, opts *v1.CreateOptions) error
	Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Comment, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.CommentList, error)
	Update(ctx context.Context, comment *v1.Comment, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error
}

type commentService struct {
	store store.Store
}

func newCommentService(srv *service) CommentService {
	return &commentService{store: srv.store}
}

func (s *commentService) Create(ctx context.Context, comment *v1.Comment, opts *v1.CreateOptions) error {
	return s.store.Comments().Create(ctx, comment, opts)
}

func (s *commentService) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Comment, error) {
	return s.store.Comments().Get(ctx, id, opts)
}

func (s *commentService) List(ctx context.Context, opts *v1.ListOptions) (*v1.CommentList, error) {
	return s.store.Comments().List(ctx, opts)
}

func (s *commentService) Update(ctx context.Context, comment *v1.Comment, opts *v1.UpdateOptions) error {
	return s.store.Comments().Update(ctx, comment, opts)
}

func (s *commentService) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.store.Comments().Delete(ctx, id, opts)
}
