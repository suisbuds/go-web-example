package repository

import (
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/app"
)


func (r *Repository) CreateUser(user *model.User) error {
	err := r.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(where interface{}) (*model.User, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user *model.User) error {
	tx := r.BeginTransaction()
	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Repository) DeleteUser(id uint32) error {
	tx:= r.BeginTransaction()
	if err := tx.Where(&model.User{Model: &model.Model{ID: id}}).Delete(&model.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r *Repository) GetUserList(page, pageSize int, where interface{}) ([]*model.User, error) {
	var users []*model.User
	offset:=app.GetPageOffset(page, pageSize)
	err := r.List(&users, where, offset, pageSize)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetUserAvatar(where interface{}) (string, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		return "", err
	}
	return user.Avatar, nil
}

func (r *Repository) GetUserID(where interface{}) (uint32, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *Repository) CheckUser(where interface{}) (bool, error) {
	var user model.User
	err := r.Get(where, &user)
	if err != nil {
		return false, err
	}
	return true, nil
}
