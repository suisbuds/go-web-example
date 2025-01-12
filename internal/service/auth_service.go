package service

import (

	"github.com/suisbuds/miao/pkg/errcode"
)

// 接口入参校验
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

// 判断认证信息 ID 是否存在
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}

	if auth.ID > 0 {
		return nil
	}

	return errcode.UnauthorizedAuthNotExist
}