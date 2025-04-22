package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Comment struct {
	ObjMeta `json:",inline"`

	Content    string `json:"content" gorm:"column:content" validate:"required"`
	Auhtor     string `json:"author" gorm:"column:author" validate:"required"`
	RefersItem uint   `json:"refer_id" gorm:"column:refers" validate:"required"`
	RefersType string `json:"refer_type" gorm:"column:refers_type" validate:"required,oneof=post problem comment"`
	SourceItem uint   `json:"source_id" gorm:"column:source" validate:"required"`
	SourceType string `json:"source_type" gorm:"column:source_type" validate:"required,oneof=post problem"`
}

type CommentList struct {
	ListMeta `json:",inline"`
	Items    []Comment `json:"items"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func (com *Comment) Validate() error {
	validator := validator.New()
	if err := validator.Struct(com); err != nil {
		return err
	}
	return nil
}

func (com *Comment) BeforeCreate(tx *gorm.DB) error {
	com.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (com *Comment) BeforeUpdate(tx *gorm.DB) error {
	com.UpdatedAt = time.Now()
	return nil
}

func (com *Comment) AfterGet(tx *gorm.DB) error {
	com.ID = 0
	return nil
}

func (com *Comment) Override(c *Comment) *Comment {
	if c.Content != "" {
		com.Content = c.Content
	}
	return com
}
