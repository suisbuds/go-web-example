package service

import (
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/app"
)

// 参数请求格式

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=1,max=50"`
	Password  string `form:"password" binding:"required,min=6"`
	Avatar    string `form:"avatar" binding:"required,url"`
	UserType  uint8  `form:"user_type" binding:"required,oneof=1 2"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"required,oneof=0 1"`
}

type GetUserRequest struct {
	Username string `form:"username" binding:"required,min=1,max=50"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetUserListRequest struct {
	UserType uint8 `form:"user_type" binding:"required,oneof=1 2"`
	State    uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	Username   string `form:"username" binding:"required,min=1,max=50"`
	Password   string `form:"password" binding:"omitempty,min=6"`
	Avatar     string `form:"avatar" binding:"omitempty,url"`
	UserType   uint8  `form:"user_type" binding:"omitempty,oneof=1 2"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
	State      uint8  `form:"state" binding:"omitempty,oneof=0 1"`
}

type DeleteUserRequest struct {
	Username string `form:"username" binding:"required,min=1,max=50"`
}

type CheckUserRequest struct {
	Username string `form:"username" binding:"required,min=1,max=50"`
	Password string `form:"password" binding:"required,min=6"`
}

type CheckUsernameRequest struct {
	Username string `form:"username" binding:"required,min=1,max=50"`
}

func (svc *Service) CreateUser(req *CreateUserRequest) error {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Avatar:   req.Avatar,
		UserType: req.UserType,
		Model:    &model.Model{CreatedBy: req.CreatedBy, State: req.State},
	}
	return svc.repo.CreateUser(user)
}

func (svc *Service) GetUser(req *GetUserRequest) (*model.User, error) {
	return svc.repo.GetUser(map[string]interface{}{"username": req.Username, "state": req.State})
}

func (svc *Service) GetUserList(req *GetUserListRequest, pager *app.Pager) ([]*model.User, int64, error) {
	return svc.repo.GetUserList(pager.Page, pager.PageSize, map[string]interface{}{"user_type": req.UserType, "state": req.State})
}

func (svc *Service) UpdateUser(req *UpdateUserRequest) error {
	values := make(map[string]interface{})
	values["modified_by"] = req.ModifiedBy
	values["username"] = req.Username
	if req.Password != "" {
		values["password"] = req.Password
	}
	if req.Avatar != "" {
		values["avatar"] = req.Avatar
	}
	if req.UserType == 1 || req.UserType == 2 {
		values["user_type"] = req.UserType
	}
	if req.State == 0 || req.State == 1 {
		values["state"] = req.State
	}
	return svc.repo.UpdateUser(values, map[string]interface{}{"username": req.Username})
}

func (svc *Service) DeleteUser(req *DeleteUserRequest) error {
	return svc.repo.DeleteUser(map[string]interface{}{"username": req.Username})
}

func (svc *Service) GetUserAvatar(req *GetUserRequest) (string, error) {
	return svc.repo.GetUserAvatar(map[string]interface{}{"username": req.Username, "state": req.State})
}

func (svc *Service) CheckUser(req *CheckUserRequest) (bool, error) {
	return svc.repo.CheckUser(map[string]interface{}{"username": req.Username, "password": req.Password})
}

func (svc *Service) CheckUsername(req *CheckUsernameRequest) (bool, error) {
	return svc.repo.CheckUser(map[string]interface{}{"username": req.Username})
}
