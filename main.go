package main

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/models"
	"github.com/suisbuds/miao/internal/routers"
	"github.com/suisbuds/miao/pkg/logger"
	"github.com/suisbuds/miao/pkg/setting"
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
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	err = setupZap()
	if err != nil {
		log.Fatalf("init.setupZap err: %v", err)
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

	global.Logger.Logf(logger.DEBUG, logger.SINGLE, "%s: miao_blog/%s", "suisbuds", "miao")
	global.Zapper.Debugf("%s: miao_blog/%s", "suisbuds", "miao")
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
	// Viper 读取配置
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

func setupZap() error {
	// zap 日志写入，lumberjack 管理日志滚动
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

func setupDBEngine() error {
	var err error

	global.DBEngine, err = models.NewDBEngine(global.DatabaseSetting)

	if err != nil {
		return err
	}

	return nil
}

