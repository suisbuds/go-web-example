package api

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/service"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/logger"
	"go.uber.org/zap/zapcore"
)

// jwt-auth API 端点

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs: %v", errs)
		global.Zapper.Logf(zapcore.ErrorLevel, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.CheckAuth err: %v", err)
		global.Zapper.Logf(zapcore.ErrorLevel, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.GenerateToken err: %v", err)
		global.Zapper.Logf(zapcore.ErrorLevel, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
