package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type postStore struct {
	db *gorm.DB
}

func newPostStore(db *gorm.DB) *postStore {
	return &postStore{db: db}
}

func (s *postStore) Create(ctx context.Context, post *v1.Post, opts *v1.CreateOptions) error {
	return s.db.Create(post).Error
}

func (s *postStore) Get(ctx context.Context, insId uint, opts *v1.GetOptions) (*v1.Post, error) {
	post := new(v1.Post)
	err := s.db.Where("instance_id", insId).First(post).Error
	return post, err
}

func (s *postStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.PostList, error) {
	var posts []v1.Post
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Post{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Model(&v1.Post{}).Offset(opts.Offset).Limit(opts.Limit).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return &v1.PostList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: posts,
	}, nil
}

func (s *postStore) Update(ctx context.Context, post *v1.Post, opts *v1.UpdateOptions) error {
	return s.db.Model(&v1.Post{}).Where("instance_id = ?", post.ObjMeta.InstanceID).Updates(post).Error
}

func (s *postStore) Delete(ctx context.Context, insId uint, opts *v1.DeleteOptions) error {
	return s.db.Where("instance_id = ?", insId).Delete(&v1.Post{}).Error
}
