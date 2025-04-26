package substore

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
)

type AwsStore interface {
	Put(ctx context.Context, opts *v1.PutOptions) (*v1.Subject, error)
}
