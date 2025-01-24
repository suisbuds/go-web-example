package repository

import (
	"gorm.io/gorm"
)

// Repository 层封装数据库操作

type Repository struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Repository {
	return &Repository{
		engine: engine,
	}
}

func (r *Repository) create(value interface{}) error {
	if err := r.engine.Create(value).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) get(where interface{}, out interface{}) error {
	if err := r.engine.Where(where).First(out).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) update(model interface{}, where interface{}, values interface{}) error {
	if err := r.engine.Model(model).Where(where).Updates(values).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) delete(model interface{}, where interface{}) error {
	if err := r.engine.Where(where).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) list(out interface{}, where interface{}, pageOffset, pageSize int) error {
	if err := r.engine.Where(where).Offset(pageOffset).Limit(pageSize).Find(out).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) count(model interface{}, where interface{}) (int64, error) {
	var count int64
	if err := r.engine.Model(model).Where(where).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 事务

func (r *Repository) save(value interface{}) error {
	if err := r.engine.Save(value).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) beginTransaction() *gorm.DB {
	return r.engine.Begin()
}
