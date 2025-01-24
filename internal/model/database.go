package model

import (
	"fmt"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type database struct {
	*gorm.DB
}

func (d *database) connect(databaseSetting *setting.DatabaseSetting) error {

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

	// gorm 连接 postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	// 获取 DB 实例
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	d.DB = db

	return nil
}

func (d *database) getDB() *gorm.DB {
	return d.DB
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	db := &database{}
	err := db.connect(databaseSetting)
	return db.getDB(), err
}
