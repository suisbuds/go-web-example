package service

import (
	"context"

	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/repository"
)

// Service 层封装具体业务逻辑, 执行参数校验

type Service struct {
	ctx  context.Context
	repo *repository.Repository
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.repo = repository.New(global.DBEngine)
	return svc
}
