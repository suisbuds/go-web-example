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
	LogFileExt           string
	UploadSavePath       string
	UploadServerUrl      string
	UploadImageMaxSize   int
	UploadImageAllowExts []string
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

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
