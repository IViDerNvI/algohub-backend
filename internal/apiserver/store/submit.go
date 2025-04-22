package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type SubmitStore interface {
	Create(ctx context.Context, submit *v1.Submit, opts *v1.CreateOptions) error
	Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Submit, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.SubmitList, error)
	Update(ctx context.Context, submit *v1.Submit, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error
}
