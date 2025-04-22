package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type PostService interface {
	Create(ctx context.Context, post *v1.Post, opts *v1.CreateOptions) error
	Get(ctx context.Context, insId uint, opts *v1.GetOptions) (*v1.Post, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.PostList, error)
	Update(ctx context.Context, post *v1.Post, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, insId uint, opts *v1.DeleteOptions) error
}

type postService struct {
	store store.Store
}

func newPostService(srv *service) PostService {
	return &postService{store: srv.store}
}

func (s *postService) Create(ctx context.Context, post *v1.Post, opts *v1.CreateOptions) error {
	return s.store.Posts().Create(ctx, post, opts)
}

func (s *postService) Get(ctx context.Context, insId uint, opts *v1.GetOptions) (*v1.Post, error) {
	return s.store.Posts().Get(ctx, insId, opts)
}

func (s *postService) List(ctx context.Context, opts *v1.ListOptions) (*v1.PostList, error) {
	return s.store.Posts().List(ctx, opts)
}

func (s *postService) Update(ctx context.Context, post *v1.Post, opts *v1.UpdateOptions) error {
	return s.store.Posts().Update(ctx, post, opts)
}

func (s *postService) Delete(ctx context.Context, insId uint, opts *v1.DeleteOptions) error {
	return s.store.Posts().Delete(ctx, insId, opts)
}
