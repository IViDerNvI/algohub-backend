package cache

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type TokenCache interface {
	Set(ctx context.Context, token string, opts *v1.CreateOptions) error
	Get(ctx context.Context, token string, opts *v1.GetOptions) error
}
