package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type SolutionStore interface {
	Create(ctx context.Context, solution *v1.Solution, opts *v1.CreateOptions) error
	Update(ctx context.Context, solution *v1.Solution, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error
	Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Solution, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SolutionList, error)
}
