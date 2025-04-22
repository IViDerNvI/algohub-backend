package mysql

import (
	"context"

	v1 "github.com/ividernvi/algohub/model/v1"
	"gorm.io/gorm"
)

type subscribeStore struct {
	db *gorm.DB
}

func newSubscribeStore(db *gorm.DB) *subscribeStore {
	return &subscribeStore{db: db}
}

func (s *subscribeStore) Create(ctx context.Context, subscribe *v1.Subscribe, opts *v1.CreateOptions) error {
	return s.db.Create(subscribe).Error
}

func (s *subscribeStore) Get(ctx context.Context, sub *v1.Subscribe, opts *v1.GetOptions) (*v1.SubscribeList, error) {
	var result v1.SubscribeList
	err := s.db.Model(&v1.Subscribe{}).Where("item_type = ? AND item_id = ?", sub.ItemType, sub.ItemID).Count(&result.TotalItems).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *subscribeStore) List(ctx context.Context, opts *v1.ListOptions) (*v1.SubscribeList, error) {
	var subscribes []v1.Subscribe
	var total int64

	query := opts.ApplyListOptions(s.db)
	err := query.Model(&v1.Subscribe{}).Count(&total).Error
	if err != nil {
		return nil, err
	}

	err = query.Offset(opts.Offset).Limit(opts.Limit).Find(&subscribes).Error
	if err != nil {
		return nil, err
	}

	return &v1.SubscribeList{
		ListMeta: v1.ListMeta{
			TotalItems: total,
		},
		Items: subscribes,
	}, nil
}

func (s *subscribeStore) Update(ctx context.Context, subscribe *v1.Subscribe, opts *v1.UpdateOptions) error {
	return s.db.Save(subscribe).Error
}

func (s *subscribeStore) Delete(ctx context.Context, sub *v1.Subscribe, opts *v1.DeleteOptions) error {
	return s.db.Where("username = ? AND item_type = ? AND item_id = ?", sub.UserName, sub.ItemType, sub.ItemID).Delete(&v1.Subscribe{}).Error
}
