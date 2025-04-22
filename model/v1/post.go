package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Post struct {
	ObjMeta `json:",inline"`

	Title   string `json:"title" gorm:"column:title" validate:"required,min=5,max=30"`
	Content string `json:"content" gorm:"column:content" validate:"required,min=10,max=500"`

	Author string `json:"author" gorm:"column:author" validate:"required"`
}

type PostList struct {
	ListMeta `json:",inline"`
	Items    []Post `json:"items"`
}

func (p *Post) TableName() string {
	return "posts"
}

func (p *Post) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func (p *Post) BeforeCreate(tx *gorm.DB) error {
	p.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (p *Post) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Post) AfterGet(tx *gorm.DB) error {
	p.ID = 0
	return nil
}

func (p *Post) Override(newpost *Post) *Post {
	if newpost.Title != "" {
		p.Title = newpost.Title
	}
	if newpost.Content != "" {
		p.Content = newpost.Content
	}
	if newpost.Author != "" {
		p.Author = newpost.Author
	}
	return p
}
