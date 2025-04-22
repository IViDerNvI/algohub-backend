package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ividernvi/algohub/pkg/util/idutil"
	"gorm.io/gorm"
)

type Problem struct {
	ObjMeta     `json:",inline"`
	Unique_ID   string  `json:"unique_id" gorm:"column:unique_id;uniqueIndex;type:varchar(255)" validate:"required"`
	Title       string  `json:"title" gorm:"column:title" validate:"required"`
	Descrition  string  `json:"descrition" gorm:"column:descrition" validate:"required"`
	Author      string  `json:"author" gorm:"column:author" validate:"required"`
	TimeLimit   int     `json:"time_limit" gorm:"column:time_limit" validate:"required"`
	MemoryLimit float64 `json:"memory_limit" gorm:"column:memory_limit" validate:"required"`
	Tag         string  `json:"tag" gorm:"column:tag" validate:"required"`
	Level       int     `json:"level" gorm:"column:level" validate:"required"`
}

type ProblemList struct {
	ListMeta `json:",inline"`
	Items    []Problem `json:"items"`
}

func (p *Problem) TableName() string {
	return "problems"
}

func (p *Problem) Validate() error {
	validator := validator.New()
	if err := validator.Struct(p); err != nil {
		return err
	}
	return nil
}

func (p *Problem) BeforeCreate(tx *gorm.DB) error {
	p.InstanceID = uint(idutil.SnowflakeID())
	return nil
}

func (p *Problem) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Problem) AfterGet(tx *gorm.DB) error {
	p.ID = 0
	return nil
}

func (p *Problem) Override(c *Problem) *Problem {
	if c.Title != "" {
		p.Title = c.Title
	}
	if c.Descrition != "" {
		p.Descrition = c.Descrition
	}
	if c.Author != "" {
		p.Author = c.Author
	}
	if c.TimeLimit != 0 {
		p.TimeLimit = c.TimeLimit
	}
	if c.MemoryLimit != 0 {
		p.MemoryLimit = c.MemoryLimit
	}
	if c.Tag != "" {
		p.Tag = c.Tag
	}
	if c.Level != 0 {
		p.Level = c.Level
	}
	return p
}
