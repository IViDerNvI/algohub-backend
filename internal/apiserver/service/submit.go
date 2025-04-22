package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type SubmitService interface {
	Create(ctx context.Context, submit *v1.Submit, opts *v1.CreateOptions) error
	Get(ctx context.Context, instance_id uint, opts *v1.GetOptions) (*v1.Submit, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SubmitList, error)
	Update(ctx context.Context, submit *v1.Submit, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, instance_id uint, opts *v1.DeleteOptions) error
}

type submitService struct {
	store store.Store
}

func newSubmitService(srv *service) SubmitService {
	return &submitService{store: srv.store}
}

func (s *submitService) Create(ctx context.Context, submit *v1.Submit, opts *v1.CreateOptions) error {

	return s.store.Submits().Create(ctx, submit, opts)
}

func (s *submitService) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Submit, error) {
	return s.store.Submits().Get(ctx, id, opts)
}

func (s *submitService) List(ctx context.Context, opts *v1.ListOptions) (*v1.SubmitList, error) {
	return s.store.Submits().List(ctx, opts)
}

func (s *submitService) Update(ctx context.Context, submit *v1.Submit, opts *v1.UpdateOptions) error {
	return s.store.Submits().Update(ctx, submit, opts)
}

func (s *submitService) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.store.Submits().Delete(ctx, id, opts)
}
