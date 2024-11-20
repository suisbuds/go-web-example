package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/models"
	"github.com/suisbuds/miao/internal/routers"
	"github.com/suisbuds/miao/pkg/setting"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = testDBConnection()
	if err != nil {
		log.Fatalf("init.testDBConnection err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("ServerSetting: %+v\n", global.ServerSetting)
	fmt.Printf("AppSetting: %+v\n", global.AppSetting)
	fmt.Printf("DatabaseSetting: %+v\n", global.DatabaseSetting)
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupLogger() error {
	return nil
}

func setupDBEngine() error {

	var err error
	// 全局变量赋值
	global.DBEngine, err = models.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

// 测试数据库连接
func testDBConnection() error {
	sqlDB, err := global.DBEngine.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Database connection successful!")
	return nil
}
