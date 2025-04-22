package model

import (
	"time"

	"gorm.io/gorm"
)

type ObjMeta struct {
	ID         uint           `gorm:"primary_key"`
	InstanceID uint           `json:"instanceID" gorm:"column:instance_id;unique"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"` // 启用逻辑删除
}

type ListMeta struct {
	TotalItems int64 `json:"totalItems"`
}
