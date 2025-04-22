package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type LikeStore interface {
	Create(ctx context.Context, like *v1.Like, opts *v1.CreateOptions) error
	Delete(ctx context.Context, like *v1.Like, opts *v1.DeleteOptions) error
	Get(ctx context.Context, like *v1.Like, opts *v1.GetOptions) (*v1.LikeList, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.LikeList, error)
	Update(ctx context.Context, like *v1.Like, opts *v1.UpdateOptions) error
}
