package repository

import (
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/logger"
)

func (r *Repository) CreateRole(role *model.Role) error {
	err := r.Create(role)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Create role failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) GetRole(where interface{}) (*model.Role, error) {
	var role model.Role
	err := r.Get(where, &role)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get role failed: %v", err)
		return nil, err
	}
	return &role, nil
}


func (r *Repository) UpdateRole(role *model.Role, where interface{}, values interface{}) error {
	err := r.Update(role, where, values)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "Update role failed: %v", err)
		return err
	}
	return nil
}

func (r *Repository) DeleteRole(id uint32) error {
    err := r.engine.Where(&model.Role{Model: &model.Model{ID: id}}).Delete(&model.Role{}).Error
    if err != nil {
        global.Logger.Logf(logger.ERROR, logger.SINGLE, "Delete role failed: %v", err)
        return err
    }
    return nil
}

func (r *Repository) GetUserRoles(userID uint32) ([]*model.Role, error) {
    var roles []*model.Role
    err := r.engine.Where("user_id = ?", userID).Find(&roles).Error
    if err != nil {
        global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get user roles failed: %v", err)
        return nil, err
    }
    return roles, nil
}

func (r *Repository) GetRoleList(pageOffset, pageSize int, where interface{}) ([]*model.Role, error) {
    var roles []*model.Role
    err := r.List(&roles, where, pageOffset, pageSize)
    if err != nil {
        global.Logger.Logf(logger.ERROR, logger.SINGLE, "Get role list failed: %v", err)
        return nil, err
    }
    return roles, nil
}
