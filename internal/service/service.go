package service

import (
	"context"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/dao"
)

// Service 层封装业务逻辑, 与 Dao 层交互, 管理数据库事务, 定义接口

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
