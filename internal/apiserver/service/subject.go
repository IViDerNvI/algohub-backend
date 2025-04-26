package service

import (
	"context"

	"github.com/ividernvi/algohub/internal/apiserver/substore"
	v1 "github.com/ividernvi/algohub/model/v1"
)

type SubjectService interface {
	Put(ctx context.Context, opts *v1.PutOptions) (*v1.Subject, error)
}

type subjectService struct {
	awsStore substore.SubStore
}

func newSubjectService(s substore.SubStore) SubjectService {
	return &subjectService{
		awsStore: s,
	}
}

func (s *subjectService) Put(ctx context.Context, opts *v1.PutOptions) (*v1.Subject, error) {
	return s.awsStore.AwsStore().Put(ctx, opts)
}
