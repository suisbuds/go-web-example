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
