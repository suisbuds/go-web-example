package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/suisbuds/miao/global"
	"github.com/suisbuds/miao/internal/service"
	"github.com/suisbuds/miao/pkg/app"
	"github.com/suisbuds/miao/pkg/errcode"
	"github.com/suisbuds/miao/pkg/logger"
)

// api controller 层处理 HTTP 路由请求, 编写接口注解

type User struct{}

func NewUser() User {
	return User{}
}

// @Summary 创建用户
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Param avatar body string true "头像"
// @Param user_type body int true "用户类型"
// @Param created_by body string true "创建者"
// @Param state body int true "状态"
// @Success 200 {object} service.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [post]
func (u User) CreateUser(c *gin.Context) {
	req := service.CreateUserRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &req)
	// 接口参数错误
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	exist, err := svc.CheckUsername(&service.CheckUsernameRequest{Username: req.Username})
	if err != nil || exist {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.CheckUser err:%v", err)
		response.ToErrorResponse(errcode.ErrorCheckUsernameFail)
		return
	}
	err = svc.CreateUser(&req)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.CreateUser err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "create user success"})
}

// @Summary 获取用户
// @Produce json
// @Param username body string true "用户名"
// @Param state body int false "状态"
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/{username} [get]
func (u User) GetUser(c *gin.Context) {
	req := service.GetUserRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.GetUser(&req)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.GetUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserFail)
		return
	}

	response.ToResponse(user)
}

// @Summary 获取用户列表
// @Produce json
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Param user_type body int true "用户类型"
// @Param state body int false "状态"
// @Success 200 {object} model.UserSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user [get]
func (u User) GetUserList(c *gin.Context) {
	req := service.GetUserListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	users, total, err := svc.GetUserList(&req, &pager)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.GetUserList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserListFail)
		return
	}

	response.ToResponseList(users, total)
}

// @Summary 更新用户
// @Produce json
// @Param username body string true "用户名"
// @Param avatar body string false "头像"
// @Param user_type body int false "用户类型"
// @Param modified_by body string true "修改者"
// @Param state body int false "状态"
// @Success 200 {object} service.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/{username} [put]
func (u User) UpdateUser(c *gin.Context) {
	req := service.UpdateUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	exist, err := svc.CheckUsername(&service.CheckUsernameRequest{Username: req.Username})
	if err != nil || exist {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.CheckUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorCheckUsernameFail)
		return
	}
	err = svc.UpdateUser(&req)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.UpdateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "update user success"})
}

// @Summary 删除用户
// @Produce json
// @Param username body string true "用户名"
// @Success 200 {object} service.Response "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/user/{username} [delete]
func (u User) DeleteUser(c *gin.Context) {
	req := service.DeleteUserRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteUser(&req)
	if err != nil {
		global.Logger.Logf(logger.ERROR, logger.SINGLE, "svc.DeleteUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteUserFail)
		return
	}

	response.ToResponse(gin.H{"message": "delete user success"})
}
