package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/setting"
)

func TestSetupSetting(t *testing.T) {
	err := setupSetting()
	assert.NoError(t, err, "setupSetting should not return an error")

	assert.Equal(t, "debug", global.ServerSetting.RunMode, "RunMode should be 'debug'")
	assert.Equal(t, "8000", global.ServerSetting.HttpPort, "HttpPort should be '8000'")
	assert.Equal(t, 60*time.Second, global.ServerSetting.ReadTimeout, "ReadTimeout should be 60 seconds")
	assert.Equal(t, 60*time.Second, global.ServerSetting.WriteTimeout, "WriteTimeout should be 60 seconds")

	assert.Equal(t, 10, global.AppSetting.DefaultPageSize, "DefaultPageSize should be 10")
	assert.Equal(t, 100, global.AppSetting.MaxPageSize, "MaxPageSize should be 100")
	assert.Equal(t, "storage/logs", global.AppSetting.LogSavePath, "LogSavePath should be 'storage/logs'")
	assert.Equal(t, "app", global.AppSetting.LogFileName, "LogFileName should be 'app'")
	assert.Equal(t, ".log", global.AppSetting.LogFileExt, "LogFileExt should be '.log'")

	assert.Equal(t, "postgres", global.DatabaseSetting.DBType, "DBType should be 'postgres'")
	assert.Equal(t, "postgres", global.DatabaseSetting.UserName, "UserName should be 'postgres'")
	assert.Equal(t, "", global.DatabaseSetting.Password, "Password should be empty")
	assert.Equal(t, "127.0.0.1", global.DatabaseSetting.Host, "Host should be '127.0.0.1'")
	assert.Equal(t, "5432", global.DatabaseSetting.Port, "Port should be '5432'")
	assert.Equal(t, "miao", global.DatabaseSetting.DBName, "DBName should be 'miao'")
	assert.Equal(t, "miao_", global.DatabaseSetting.TablePrefix, "TablePrefix should be 'miao_'")
	assert.Equal(t, "disable", global.DatabaseSetting.SSLMode, "SSLMode should be 'disable'")
	assert.Equal(t, "Asia/Shanghai", global.DatabaseSetting.TimeZone, "TimeZone should be 'Asia/Shanghai'")
	assert.Equal(t, 10, global.DatabaseSetting.MaxIdleConns, "MaxIdleConns should be 10")
	assert.Equal(t, 30, global.DatabaseSetting.MaxOpenConns, "MaxOpenConns should be 30")

	assert.Equal(t, global.ServerSetting.ReadTimeout, 60*time.Second, "ReadTimeout should be 60")
	assert.Equal(t, global.ServerSetting.WriteTimeout, 60*time.Second, "WriteTimeout should be 60")
}

func TestSetupDBEngine(t *testing.T) {
	global.DatabaseSetting = &setting.DatabaseSetting{
		DBType:       "postgres",
		UserName:     "postgres",
		Password:     "",
		Host:         "127.0.0.1",
		Port:         "5432",
		DBName:       "miao",
		TablePrefix:  "miao_",
		SSLMode:      "disable",
		TimeZone:     "Asia/Shanghai",
		MaxIdleConns: 10,
		MaxOpenConns: 30,
	}

	err := setupDBEngine()
	assert.NoError(t, err, "setupDBEngine should not return an error")

	// 测试数据库连接
	err = global.DBEngine.Exec("SELECT 1").Error
	assert.NoError(t, err, "Database connection should be successful")
}
