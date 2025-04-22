package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Like struct {
	ObjMeta  `json:",inline"`
	UserName string `json:"username" gorm:"primaryKey;column:username"`
	ItemType string `json:"item_type" gorm:"primaryKey;column:item_type" validate:"required;oneof:post,comment"`
	ItemID   uint   `json:"item_id" gorm:"primaryKey;column:item_id" validate:"required"`
}

type LikeList struct {
	ListMeta `json:",inline"`
	Items    []Like `json:"items"`
}

func (l *Like) TableName() string {
	return "likes"
}

func (l *Like) Validate() error {
	validator := validator.New()
	if err := validator.Struct(l); err != nil {
		return err
	}
	return nil
}

func (l *Like) BeforeCreate(tx *gorm.DB) error {
	l.InstanceID = uint(idutil.SnowflakeID())
	return nil
}
