package repository

import (
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/app"
)

func (r *Repository) CreateUser(user *model.User) error {
	err := r.create(user)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUser(where interface{}) (*model.User, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) UpdateUser(values interface{}, where interface{}) error {
	if err := r.update(&model.User{}, where, values); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteUser(id uint32) error {
	if err := r.delete(&model.User{}, &model.User{Model: &model.Model{ID: id}}); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserList(page, pageSize int, where interface{}) ([]*model.User, int64, error) {
	var users []*model.User
	offset := app.GetPageOffset(page, pageSize)
	err := r.list(&users, where, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := r.count(&model.User{}, where)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

func (r *Repository) GetUserAvatar(where interface{}) (string, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return "", err
	}
	return user.Avatar, nil
}

func (r *Repository) GetUserID(where interface{}) (uint32, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *Repository) CheckUser(where interface{}) (bool, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return false, err
	}
	return true, nil
}
