package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type ProblemService interface {
	Create(ctx context.Context, problem *v1.Problem, opts *v1.CreateOptions) error
	Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Problem, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.ProblemList, error)
	Update(ctx context.Context, problem *v1.Problem, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error
}

type problemService struct {
	store store.Store
}

func newProblemService(srv *service) ProblemService {
	return &problemService{store: srv.store}
}

func (s *problemService) Create(ctx context.Context, problem *v1.Problem, opts *v1.CreateOptions) error {
	return s.store.Problems().Create(ctx, problem, opts)
}

func (s *problemService) Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Problem, error) {
	return s.store.Problems().Get(ctx, name, opts)
}

func (s *problemService) List(ctx context.Context, opts *v1.ListOptions) (*v1.ProblemList, error) {
	return s.store.Problems().List(ctx, opts)
}

func (s *problemService) Update(ctx context.Context, problem *v1.Problem, opts *v1.UpdateOptions) error {
	return s.store.Problems().Update(ctx, problem, opts)
}

func (s *problemService) Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error {
	return s.store.Problems().Delete(ctx, name, opts)
}
