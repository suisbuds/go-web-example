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


type Database struct {
	DB *gorm.DB
}

func (db *Database) connect(*setting.DatabaseSetting) error {

	// 构建数据源 DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		global.DatabaseSetting.Host,
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.DBName,
		global.DatabaseSetting.Port,
		global.DatabaseSetting.SSLMode,
		global.DatabaseSetting.TimeZone)

	// 使用 gorm 默认 logger
	var newLogger logger.Interface
	if global.ServerSetting.RunMode == "debug" {
		newLogger = logger.Default.LogMode(logger.Info)
	} else {
		newLogger = logger.Default.LogMode(logger.Silent)
	}

	// gorm 连接 pg
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	// 获取 DB 对象
	sqlDB, err := database.DB()
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(global.DatabaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.DatabaseSetting.MaxOpenConns)

	db.DB = database

	return nil
}

func (db *Database) getDB() *gorm.DB {
	return db.DB
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {
	db := &Database{}
	err:=db.connect(databaseSetting)
	return db.getDB(), err
}
