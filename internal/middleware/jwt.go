package middleware

import (
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/internal/service"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/logger"
)

func InitJWT() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "miao", // 中间件名称
		Key:         []byte(global.JWTSetting.Secret),
		Timeout:     global.JWTSetting.Timeout, // token 过期时间
		MaxRefresh:  global.JWTSetting.MaxRefresh,
		PayloadFunc: payloadFunc(), // 登陆回调

		Authenticator: authenticator(),                                    // 根据登陆信息验证用户
		Authorizator:  authorizator(),                                     // 控制用户权限
		Unauthorized:  unauthorized(),                                     // 未授权处理
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // token 检索模式
		TokenHeadName: "suisbuds",                                         // token 请求头名称
		TimeFunc:      time.Now,                                           // 时区
	}
}

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.User); ok {
			return jwt.MapClaims{
				"username": v.Username,
			}
		}
		return jwt.MapClaims{}
	}
}

func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		req := service.CheckUserRequest{}
		response := app.NewResponse(c)
		valid, errs := app.BindAndValid(c, &req)
		if !valid {
			global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs:%v", errs)
			response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
			return "", jwt.ErrFailedAuthentication
		}
		svc := service.New(c.Request.Context())
		ok, err := svc.CheckUser(&req)
		if err != nil || !ok {
			global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.CheckUser err:%v", err)
			response.ToErrorResponse(errcode.ErrorCheckUserFail)
			return "", jwt.ErrFailedAuthentication
		}
		loginUser := model.User{Username: req.Username, Password: req.Password}
		return &loginUser, nil
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*model.User); ok && v.UserType == 1 {
			return true
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}
