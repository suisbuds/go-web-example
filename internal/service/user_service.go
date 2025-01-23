package service

import (
	"github.com/suisbuds/miao/internal/model"
	"github.com/suisbuds/miao/pkg/app"
)

// 参数请求格式

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=3,max=50"`
	Password  string `form:"password" binding:"required,min=6"`
	Avatar    string `form:"avatar" binding:"omitempty,url"`
	UserType  uint8  `form:"user_type,default=2" binding:"oneof=1 2"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type GetUserRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type GetUserListRequest struct {
	UserType uint8 `form:"user_type" binding:"omitempty,oneof=1 2"`
	State    uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateUserRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Username   string `form:"username" binding:"omitempty,min=3,max=50"`
	Avatar     string `form:"avatar" binding:"omitempty,url"`
	UserType   uint8  `form:"user_type" binding:"omitempty,oneof=1 2"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CheckUserRequest struct {
	Username string `form:"username" binding:"required,min=3,max=50"`
	Password string `form:"password" binding:"required,min=6"`
}

func (svc *Service) CreateUser(req *CreateUserRequest) error {
	user := &model.User{
		Username: req.Username,
		Password: req.Password,
		Avatar:   req.Avatar,
		UserType: req.UserType,
	}
	return svc.repo.CreateUser(user)
}

func (svc *Service) GetUser(req *GetUserRequest) (*model.User, error) {
	return svc.repo.GetUser(map[string]interface{}{"id": req.ID, "state": req.State})
}

func (svc *Service) GetUserList(req *GetUserListRequest, pager *app.Pager) ([]*model.User, error) {
	return svc.repo.GetUserList(pager.Page, pager.PageSize, map[string]interface{}{"user_type": req.UserType, "state": req.State})
}

func (svc *Service) UpdateUser(req *UpdateUserRequest) error {
	values := make(map[string]interface{})
	values["modified_by"] = req.ModifiedBy
	if req.Username != "" {
		values["username"] = req.Username
	}
	if req.Avatar != "" {
		values["avatar"] = req.Avatar
	}
	if req.UserType == 1 || req.UserType == 2 {
		values["user_type"] = req.UserType
	}
	return svc.repo.UpdateUser(values, map[string]interface{}{"id": req.ID})
}

func (svc *Service) DeleteUser(req *DeleteUserRequest) error {
	return svc.repo.DeleteUser(req.ID)
}

func (svc *Service) GetUserAvatar(req *GetUserRequest) (string, error) {
	return svc.repo.GetUserAvatar(&model.User{Model: &model.Model{ID: req.ID}})
}

func (svc *Service) CheckUser(req *CheckUserRequest) (bool, error) {
	return svc.repo.CheckUser(&model.User{Username: req.Username, Password: req.Password})
}
