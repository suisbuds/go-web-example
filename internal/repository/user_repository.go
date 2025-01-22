package repository

import (
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/logger"
)


func (r *Repository) CreateUser(user *model.User) error {
	err := r.Create(user)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Create user failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) GetUser(where interface{}) (*model.User, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get user failed: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user *model.User, role *model.Role) error {
	tx := r.BeginTransaction()
	if err := tx.Save(user).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Update user failed: %v", err)
		tx.Rollback()
		return err
	}
	if err := tx.Save(role).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Update role failed: %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Repository) DeleteUser(id uint32) error {
	tx:= r.BeginTransaction()
	if err := tx.Where(&model.Role{UserID: id}).Delete(&model.Role{}).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Delete role failed: %v", err)
		tx.Rollback()
		return err
	}
	if err := tx.Where(&model.User{Model: &model.Model{ID: id}}).Delete(&model.User{}).Error; err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Delete user failed: %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Repository) GetUserList(pageOffset, pageSize int, where interface{}) ([]*model.User, error) {
	var users []*model.User
	err := r.List(&users, where, pageOffset, pageSize)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get user list failed: %v", err)
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserAvatar(where interface{}) (string, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get user avatar failed: %v", err)
		return "", err
	}
	return user.Avatar, nil
}

func (r *Repository) GetUserID(where interface{}) (uint32, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get user ID failed: %v", err)
		return 0, err
	}
	return user.ID, nil
}

func (r *Repository) CheckUser(where interface{}) (bool, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Check user failed: %v", err)
		return false, err
	}
	return true, nil
}
