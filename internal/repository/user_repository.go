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

func (r *Repository) UpdateUser(values interface{}, where interface{}) error {
	if err := r.update(&model.User{}, where, values); err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteUser(where interface{}) error {
	if err := r.delete(&model.User{}, where); err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserAvatar(where interface{}) (string, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return "", err
	}
	return user.Avatar, nil
}

func (r *Repository) CheckUser(where interface{}) (bool, error) {
	var user model.User
	err := r.get(where, &user)
	if err != nil {
		return false, err
	}
	return true, nil
}
