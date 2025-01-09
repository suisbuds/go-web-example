package models

import (
	"fmt"
	"time"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Model 层定义数据结构, 实现 ORM 映射, 定义字段验证规则和约束, 与数据库中的表一一对应

// 公共字段设置为 CreatedAt, ModifiedAt, DeletedAt 时 Gorm 可以自动更新
// IsDel 软删除标记, ID 主键, State 状态标记, 表示是否启用资源
type Model struct {
	ID         uint32         `gorm:"primary_key" json:"id"`
	CreatedBy  string         `json:"created_by"`
	ModifiedBy string         `json:"modified_by"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	IsDel      uint8          `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {

	// 构建数据源 DSN
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.DBName,
		databaseSetting.Port,
		databaseSetting.SSLMode,
		databaseSetting.TimeZone)

	// 使用 gorm 默认 logger
	var newLogger logger.Interface
	if global.ServerSetting.RunMode == "debug" {
		newLogger = logger.Default.LogMode(logger.Info)
	} else {
		newLogger = logger.Default.LogMode(logger.Silent)
	}

	// gorm 连接 pg
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	// 获取 DB 对象
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)


	return db, nil
}


// Hooks: CreatedAt, UpdatedAt, DeletedAt (Grom 内部的软删除) 会自动更新, IsDel 自定义不能自动更新
// func (m *Model) BeforeCreate(db *gorm.DB) (err error) {
// 	// 创建记录前设置时间戳
// 	m.CreatedAt = time.Now()
// 	m.UpdatedAt = time.Now()
// 	return
// }

// func (m *Model) BeforeUpdate(db *gorm.DB) (err error) {
// 	// 更新记录前更新时间戳
// 	m.UpdatedAt = time.Now()
// 	return
// }

// func (m *Model) BeforeDelete(db *gorm.DB) (err error) {
// 	// 删除记录前进行软删除
// 	m.DeletedAt.Time = time.Now()
// 	m.DeletedAt.Valid = true
// 	m.IsDel = 1
// 	return
// }
