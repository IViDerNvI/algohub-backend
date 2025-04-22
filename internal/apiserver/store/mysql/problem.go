package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type problemStore struct {
	db *gorm.DB
}

func newProblemStore(db *gorm.DB) *problemStore {
	return &problemStore{db: db}
}

func (s *problemStore) Create(ctx context.Context, problem *v1.Problem, opts *v1.CreateOptions) error {
	return s.db.Create(problem).Error
}

func (s *problemStore) Get(ctx context.Context, name string, opts *v1.GetOptions) (*v1.Problem, error) {
	problem := new(v1.Problem)
	err := s.db.Where("unique_id = ?", name).First(problem).Error
	return problem, err
}

func (s *problemStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.ProblemList, error) {
	var problem []v1.Problem
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Problem{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&problem).Error
	if err != nil {
		return nil, err
	}

	return &v1.ProblemList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: problem,
	}, nil
}

func (s *problemStore) Update(ctx context.Context, problem *v1.Problem, opts *v1.UpdateOptions) error {
	return s.db.Updates(problem).Error
}

func (s *problemStore) Delete(ctx context.Context, name string, opts *v1.DeleteOptions) error {
	return s.db.Where("unique_id = ?", name).Delete(&v1.Problem{}).Error
}
