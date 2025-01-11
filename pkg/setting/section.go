package setting

// 利用 Viper 读取配置 config.yaml

import (
	"os"
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

type JWTSettingS struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {

	err := s.viper.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	// fmt.Printf("v's type: %T\n", v)

	// 配置 doppler 环境变量. 空接口类型 interface{}, 可以持有任何值, 在运行时需要通过反射或类型断言转换为具体类型

	if dbSetting, ok := v.(**DatabaseSetting); ok {
		if (*dbSetting).Password == "" {
			(*dbSetting).Password = os.Getenv("DB_PASSWORD") // 解引用二重指针
		}
		if (*dbSetting).UserName == "" {
			(*dbSetting).UserName = os.Getenv("USERNAME")
		}
	}

	if jwtSetting, ok := v.(**JWTSettingS); ok {
		if (*jwtSetting).Secret == "" {
			(*jwtSetting).Secret = os.Getenv("SECRET")
		}
		if (*jwtSetting).Issuer == "" {
			(*jwtSetting).Issuer = os.Getenv("ISSUER")
		}
	}

	return nil
}
