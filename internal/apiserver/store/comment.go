package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type CommentStore interface {
	Create(ctx context.Context, comment *v1.Comment, opts *v1.CreateOptions) error
	Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Comment, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.CommentList, error)
	Update(ctx context.Context, comment *v1.Comment, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error
}
