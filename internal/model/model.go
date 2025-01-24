package model

import (
	"time"

	"gorm.io/gorm"
)

// Model 层定义实体和字段标签,  与数据库中的表一一对应

// Gorm 可以自动更新公共字段 CreatedAt, ModifiedAt, DeletedAt
type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	State      uint8          `gorm:"not null;default:1" json:"state"`
}


