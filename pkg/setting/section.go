package setting

// 利用 Viper 读取配置 config.yaml

import (
	"time"
)

type ServerSetting struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSetting struct {
	DefaultPageSize      int
	MaxPageSize          int
	LogSavePath          string
	LoggerFileName       string
	ZapperFileName       string
	AccesserFileName     string
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
	ContextTimeout       time.Duration
}

type DatabaseSetting struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	TablePrefix  string
	SSLMode      string
	TimeZone     string
	MaxIdleConns int
	MaxOpenConns int
}

type JWTSetting struct {
	Secret      string
	Issuer      string
	Timeout     time.Duration
	MaxRefresh  time.Duration
	Realm	   string
}

type EmailSetting struct {
	Host     string
	Port     int
	UserName string
	Password string
	IsSSL    bool
	From     string
	To       []string
}

func (s *Setting) ReadSection(k string, v interface{}) error {

	err := s.viper.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	// 配置 doppler 环境变量. 空接口类型 interface{}, 可以持有任何值, 在运行时需要通过反射或类型断言转换为具体类型

	SetEnv(v)

	return nil
}
