package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type submitStore struct {
	db *gorm.DB
}

func newSubmitStore(db *gorm.DB) *submitStore {
	return &submitStore{db: db}
}

func (s *submitStore) Create(ctx context.Context, submit *v1.Submit, opts *v1.CreateOptions) error {
	return s.db.Create(submit).Error
}

func (s *submitStore) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Submit, error) {
	submit := new(v1.Submit)
	err := s.db.Where("instance_id = ?", id).First(submit).Error
	return submit, err
}

func (s *submitStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.SubmitList, error) {
	var submits []v1.Submit
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Submit{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&submits).Error
	if err != nil {
		return nil, err
	}

	return &v1.SubmitList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: submits,
	}, nil
}

func (s *submitStore) Update(ctx context.Context, submit *v1.Submit, opts *v1.UpdateOptions) error {
	return s.db.Save(submit).Error
}

func (s *submitStore) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.db.Where("instance_id = ?", id).Delete(&v1.Submit{}).Error
}
