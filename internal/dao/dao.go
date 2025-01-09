package dao

import "gorm.io/gorm"

// Dao 层 (Data Access Object) 封装数据访问操作, 与 Service 层和 Model 层直接交互, 执行 CRUD

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine: engine}
}
