package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type PostStore interface {
	Create(ctx context.Context, post *v1.Post, opts *v1.CreateOptions) error
	Get(ctx context.Context, insId uint, opts *v1.GetOptions) (*v1.Post, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.PostList, error)
	Update(ctx context.Context, post *v1.Post, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, insId uint, opts *v1.DeleteOptions) error
}
