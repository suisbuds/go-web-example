package global

import (
	"github.com/suisbuds/miao/pkg/logger"
	"github.com/suisbuds/miao/pkg/setting"
	"github.com/suisbuds/miao/pkg/validator"
	"go.uber.org/zap"
	"gorm.io/gorm"

)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
	Zapper          *zap.SugaredLogger
	Accesser        *logger.Logger

	DBEngine *gorm.DB

	JWTSetting *setting.JWTSetting
	
	EmailSetting *setting.EmailSetting

	Validator *validator.MiaoValidator

)
