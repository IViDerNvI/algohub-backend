package objstore

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type AvatorStore interface {
	Get(ctx context.Context, userID string, opts *v1.GetOptions) ([]byte, error)
	Put(ctx context.Context, userID string, data []byte, opts *v1.UpdateOptions) error
}
