package setting

import "github.com/spf13/viper"

type Setting struct {
	viper *viper.Viper
}

// Viper 读取配置文件 configs/config.yaml
func NewSetting() (*Setting, error) {

	CheckEnv()

	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
