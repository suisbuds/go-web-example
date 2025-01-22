package repository

import (
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/pkg/logger"
	"gorm.io/gorm"
)

// Repository 封装数据访问操作

type Repository struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Repository {
	return &Repository{
		engine: engine,
	}
}

func (r *Repository) Create(value interface{}) error {
	if err := r.engine.Create(value).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Create model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) Get(where interface{}, out interface{}) error {
	if err := r.engine.Where(where).First(out).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) Update(model interface{}, where interface{}, values interface{}) error {
	if err := r.engine.Model(model).Where(where).Updates(values).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Update model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) Delete(model interface{}, where interface{}) error {
	if err := r.engine.Where(where).Delete(model).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Delete model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) List(out interface{}, where interface{}, pageOffset, pageSize int) error {
	if err := r.engine.Where(where).Offset(pageOffset).Limit(pageSize).Find(out).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "List model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) Count(model interface{}, where interface{}) (int64, error) {
	var count int64
	if err := r.engine.Model(model).Where(where).Count(&count).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Count model failed: %v", err)
		return 0, err
	}
	return count, nil
}

func (r *Repository) Save(value interface{}) error {
	if err := r.engine.Save(value).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Save model failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) BeginTransaction() *gorm.DB {
	return r.engine.Begin()
}
