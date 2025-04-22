package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type likeStore struct {
	db *gorm.DB
}

func newLikeStore(db *gorm.DB) *likeStore {
	return &likeStore{db: db}
}

func (s *likeStore) Create(ctx context.Context, like *v1.Like, opts *v1.CreateOptions) error {
	return s.db.Create(like).Error
}

func (s *likeStore) Get(ctx context.Context, like *v1.Like, opts *v1.GetOptions) (*v1.LikeList, error) {
	var result v1.LikeList
	err := s.db.Model(&v1.Like{}).Where("item_type = ? AND item_id = ?", like.ItemType, like.ItemID).Count(&result.TotalItems).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *likeStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.LikeList, error) {
	var likes []v1.Like
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Like{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&likes).Error
	if err != nil {
		return nil, err
	}

	return &v1.LikeList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: likes,
	}, nil
}

func (s *likeStore) Update(ctx context.Context, like *v1.Like, opts *v1.UpdateOptions) error {
	return s.db.Save(like).Error
}

func (s *likeStore) Delete(ctx context.Context, like *v1.Like, opts *v1.DeleteOptions) error {
	return s.db.Where("username = ? AND item_type = ? AND item_id = ?", like.UserName, like.ItemType, like.ItemID).Delete(&v1.Like{}).Error
}
