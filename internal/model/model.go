package model

import (
	"time"

	"gorm.io/gorm"
)

// Model 层定义实体和字段标签,  与数据库中的表一一对应

// ID, CreatedAt, ModifiedAt, DeletedAt 由 gorm 自动管理
type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	State      uint8          `gorm:"not null;default:1" json:"state"`
}


