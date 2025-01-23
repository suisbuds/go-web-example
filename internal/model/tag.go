package model

import (
	"github.com/suisbuds/miao/pkg/app"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
}

// 实现 Gorm 接口, 指定 Tag 对应的数据库表名
func (t Tag) TableName() string {
	return "mio_tag"
}

// Swagger 文档生成
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}


