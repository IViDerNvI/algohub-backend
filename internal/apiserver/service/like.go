package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type LikeService interface {
	Create(ctx context.Context, like *v1.Like, opts *v1.CreateOptions) error
	Delete(ctx context.Context, like *v1.Like, opts *v1.DeleteOptions) error
	Get(ctx context.Context, like *v1.Like, opts *v1.GetOptions) (*v1.LikeList, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.LikeList, error)
	Update(ctx context.Context, like *v1.Like, opts *v1.UpdateOptions) error
}

type likeService struct {
	store store.Store
}

func newLikeService(srv *service) LikeService {
	return &likeService{store: srv.store}
}

func (s *likeService) Create(ctx context.Context, like *v1.Like, opts *v1.CreateOptions) error {
	return s.store.Likes().Create(ctx, like, opts)
}

func (s *likeService) Delete(ctx context.Context, like *v1.Like, opts *v1.DeleteOptions) error {
	return s.store.Likes().Delete(ctx, like, opts)
}

func (s *likeService) Get(ctx context.Context, like *v1.Like, opts *v1.GetOptions) (*v1.LikeList, error) {
	return s.store.Likes().Get(ctx, like, opts)
}

func (s *likeService) List(ctx context.Context, opts *v1.ListOptions) (*v1.LikeList, error) {
	return s.store.Likes().List(ctx, opts)
}

func (s *likeService) Update(ctx context.Context, like *v1.Like, opts *v1.UpdateOptions) error {
	return s.store.Likes().Update(ctx, like, opts)
}
