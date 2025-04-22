package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type commentStore struct {
	db *gorm.DB
}

func newCommentStore(db *gorm.DB) *commentStore {
	return &commentStore{db: db}
}

func (s *commentStore) Create(ctx context.Context, comment *v1.Comment, opts *v1.CreateOptions) error {
	return s.db.Create(comment).Error
}

func (s *commentStore) Get(ctx context.Context, id uint, opts *v1.GetOptions) (*v1.Comment, error) {
	comment := new(v1.Comment)
	err := s.db.Where("instance_id = ?", id).First(comment).Error
	return comment, err
}

func (s *commentStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.CommentList, error) {
	var comments []v1.Comment
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Comment{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return &v1.CommentList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: comments,
	}, nil
}

func (s *commentStore) Update(ctx context.Context, comment *v1.Comment, opts *v1.UpdateOptions) error {
	return s.db.Updates(comment).Error
}

func (s *commentStore) Delete(ctx context.Context, id uint, opts *v1.DeleteOptions) error {
	return s.db.Where("instance_id = ?", id).Delete(&v1.Comment{}).Error
}
