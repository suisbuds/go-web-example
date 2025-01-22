package service

import (
)

// 接口参数校验

type CreateUserRequest struct {
	Username  string `form:"username" binding:"required,min=3,max=50"`
	Password  string `form:"password" binding:"required,min=6"`
	Avatar    string `form:"avatar" binding:"omitempty,url"`
	UserType  uint8  `form:"user_type" binding:"oneof=1 2 3"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type GetUsersRequest struct {
	Page     int    `form:"page" binding:"required,gte=1"`
	PageSize int    `form:"page_size" binding:"required,gte=1,lte=100"`
	Username string `form:"username" binding:"omitempty,min=3,max=50"`
	UserType uint8  `form:"user_type" binding:"omitempty,oneof=1 2 3"`
}

type GetRolesRequest struct {
	UserID uint32 `form:"user_id" binding:"required,gte=1"`
}

type UpdateUserRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Username   string `form:"username" binding:"omitempty,min=3,max=50"`
	Avatar     string `form:"avatar" binding:"omitempty,url"`
	UserType   uint8  `form:"user_type" binding:"omitempty,oneof=1 2 3"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteUserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type GetUserAvatarRequest struct {
	UserID uint32 `form:"user_id" binding:"required,gte=1"`
}

type CheckUserRequest struct {
	Username string `form:"username" binding:"required,min=3,max=50"`
}


func CreateUser(){}

func UpdateUser(){}

func DeleteUser(){}

func GetUser(){}

func GetUsers(){}

func GetUserAvatar(){}

func GetRole(){}



func CheckUser(){}





