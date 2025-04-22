package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type solutionStore struct {
	db *gorm.DB
}

func newSolutionStore(db *gorm.DB) *solutionStore {
	return &solutionStore{db: db}
}

func (s *solutionStore) Create(ctx context.Context, solution *v1.Solution, opts *v1.CreateOptions) error {
	return s.db.Create(solution).Error
}

func (s *solutionStore) Update(ctx context.Context, solution *v1.Solution, opts *v1.UpdateOptions) error {
	return s.db.Save(solution).Error
}

func (s *solutionStore) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.db.Where("instance_id = ?", id).Delete(&v1.Solution{}).Error
}

func (s *solutionStore) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Solution, error) {
	var solution v1.Solution
	err := s.db.Where("instance_id = ?", id).First(&solution).Error
	if err != nil {
		return nil, err
	}
	return &solution, nil
}

func (s *solutionStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.SolutionList, error) {
	var solution []v1.Solution
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Solution{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&solution).Error
	if err != nil {
		return nil, err
	}

	return &v1.SolutionList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: solution,
	}, nil
}
