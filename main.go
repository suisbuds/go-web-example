package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/internal/router"
	"github.com/suisbuds/miao/pkg/logger"
	"github.com/suisbuds/miao/pkg/setting"
	"github.com/suisbuds/miao/pkg/validator"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
	err = setupZapper()
	if err != nil {
		log.Fatalf("init.setupZap err: %v", err)
	}
	err = setupAccesser()
	if err != nil {
		log.Fatalf("init.setupAccesser err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}
}

// @title miao
// @version 1.0
// @description My own Miao Blog
// @termsOfService https://github.com/suisbuds/miao
func main() {

	gin.SetMode(global.ServerSetting.RunMode)
	router := router.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func setupSetting() error {

	// Viper 读取配置
	s, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.AppSetting.ContextTimeout *= time.Second

	return nil
}

func setupLogger() error {

	// Viper 读取日志配置
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LoggerFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
			Compress:  true,
		},
		"",
		log.LstdFlags,
	)

	return nil
}

func setupZapper() error {

	// zap 日志写入, lumberjack 管理日志滚动
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.ZapperFileName + global.AppSetting.LogFileExt,
		MaxSize:  600,
		MaxAge:   10,
		Compress: true,
	})
	// zap 编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// zap 核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		zap.DebugLevel,
	)
	// 创建日志记录器
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(0))
	// 刷新日志缓冲区
	defer logger.Sync()

	// 转换为 SugaredLogger 并赋值给 Zapper
	global.Zapper = logger.Sugar()
	return nil
}

func setupAccesser() error {

	global.Accesser = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.AccesserFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,   // 文件存在最大天数
			LocalTime: true, // 本地时间
			Compress:  true, // 压缩
		},
		"",
		log.LstdFlags,
	)

	return nil
}

func setupDBEngine() error {

	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

func setupValidator() error {
	global.Validator = validator.NewMiaoValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator // 实现 bind.Validation 接口, 替换成 MiaoValidator
	return nil
}
