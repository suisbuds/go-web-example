package models

import (
	"fmt"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// public model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(databaseSetting *setting.DatabaseSetting) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		databaseSetting.Host,
		databaseSetting.UserName,
		databaseSetting.DBName,
		databaseSetting.Port,
		databaseSetting.SSLMode,
		databaseSetting.TimeZone)

	var logMode logger.Interface
	if global.ServerSetting.RunMode == "debug" {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logMode})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil

}
