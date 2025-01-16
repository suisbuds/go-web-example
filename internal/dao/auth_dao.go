package dao

import "github.com/suisbuds/miao/internal/models"

// 获取鉴权信息
func (d *Dao) GetAuth(appKey, appSecret string) (models.Auth, error) {
	auth := models.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
