package model

import (
	"github.com/suisbuds/miao/pkg/app"
)

// 封装 Tag 模块对应的数据操作, 与 Dao 层交互, 借助 Gorm 操作数据库
// 显式错误处理, 让代码更具可读性

type Tag struct {
	*Model
	Name  string `json:"name"`
}

// 实现 Gorm 的接口，指定 Tag 对应的 pg 表名
func (t Tag) TableName() string {
	return "mio_tag"
}

// Swagger 文档生成
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}


