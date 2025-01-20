package model

import (
	"time"

	"gorm.io/gorm"
)

// Model 层定义数据结构, 实现 ORM 映射, 定义字段验证规则和约束, 与数据库中的表一一对应

// 公共字段设置为 CreatedAt, ModifiedAt, DeletedAt 时 Gorm 可以自动更新, IsDel 软删除标记, ID 主键
type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	IsDel      uint8          `json:"is_del"`
}


