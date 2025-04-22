package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type SubscribeService interface {
	Create(ctx context.Context, subscribe *v1.Subscribe, opts *v1.CreateOptions) error
	Delete(ctx context.Context, sub *v1.Subscribe, opts *v1.DeleteOptions) error
	Get(ctx context.Context, sub *v1.Subscribe, opts *v1.GetOptions) (*v1.SubscribeList, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SubscribeList, error)
	Update(ctx context.Context, subscribe *v1.Subscribe, opts *v1.UpdateOptions) error
}

type subscribeService struct {
	store store.Store
}

func newSubscribeService(srv *service) SubscribeService {
	return &subscribeService{store: srv.store}
}

func (s *subscribeService) Create(ctx context.Context, subscribe *v1.Subscribe, opts *v1.CreateOptions) error {
	return s.store.Subscribes().Create(ctx, subscribe, opts)
}

func (s *subscribeService) Delete(ctx context.Context, sub *v1.Subscribe, opts *v1.DeleteOptions) error {
	return s.store.Subscribes().Delete(ctx, sub, opts)
}

func (s *subscribeService) Get(ctx context.Context, sub *v1.Subscribe, opts *v1.GetOptions) (*v1.SubscribeList, error) {
	return s.store.Subscribes().Get(ctx, sub, opts)
}

func (s *subscribeService) List(ctx context.Context, opts *v1.ListOptions) (*v1.SubscribeList, error) {
	return s.store.Subscribes().List(ctx, opts)
}

func (s *subscribeService) Update(ctx context.Context, subscribe *v1.Subscribe, opts *v1.UpdateOptions) error {
	return s.store.Subscribes().Update(ctx, subscribe, opts)
}
