package service

import (
	"context"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/repository"
)

// Service 层封装业务逻辑, 与 Dao 层交互, 管理数据库事务, 定义接口

type Service struct {
	ctx context.Context
	dao *repository.Repository
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = repository.New(global.DBEngine)
	return svc
}
