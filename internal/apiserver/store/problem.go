package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type ProblemStore interface {
	Create(ctx context.Context, problem *v1.Problem, opts *v1.CreateOptions) error
	Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Problem, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.ProblemList, error)
	Update(ctx context.Context, problem *v1.Problem, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error
}
