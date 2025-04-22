package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/store"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type SolutionService interface {
	Create(ctx context.Context, solution *v1.Solution, opts *v1.CreateOptions) error
	Update(ctx context.Context, solution *v1.Solution, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error
	Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Solution, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SolutionList, error)
}

type solutionService struct {
	store store.Store
}

func newSolutionService(s *service) *solutionService {
	return &solutionService{
		store: s.store,
	}
}

func (s *solutionService) Create(ctx context.Context, solution *v1.Solution, opts *v1.CreateOptions) error {
	return s.store.Solutions().Create(ctx, solution, opts)
}

func (s *solutionService) Update(ctx context.Context, solution *v1.Solution, opts *v1.UpdateOptions) error {
	return s.store.Solutions().Update(ctx, solution, opts)
}

func (s *solutionService) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.store.Solutions().Delete(ctx, id, opts)
}

func (s *solutionService) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Solution, error) {
	return s.store.Solutions().Get(ctx, id, opts)
}

func (s *solutionService) List(ctx context.Context, opts *v1.ListOptions) (*v1.SolutionList, error) {
	return s.store.Solutions().List(ctx, opts)
}
