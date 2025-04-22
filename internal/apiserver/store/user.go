package store

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type UserStore interface {
	Create(ctx context.Context, user *v1.User, opts *v1.CreateOptions) error
	Get(ctx context.Context, username string, opts *v1.GetOptions) (*v1.User, error)
	List(ctx context.Context, opts *v1.ListOptions) (*v1.UserList, error)
	Update(ctx context.Context, user *v1.User, opts *v1.UpdateOptions) error
	Delete(ctx context.Context, username string, opts *v1.DeleteOptions) error
	Verify(ctx context.Context, token string, opts *v1.VerifyOptions) (*v1.User, error)
}
