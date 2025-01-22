package repository

import "github.com/suisbuds/miao/internal/model"

// 获取鉴权信息
func (d *Repository) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
