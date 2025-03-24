package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
)

type UserApi struct {
}

// SendCode 发送验证码
func (api *UserApi) SendCode(ctx *gin.Context) {
	var req user.SendCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	resp, err := rpc.UserClient.SendCode(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// Login 用户登录
func (api *UserApi) Login(ctx *gin.Context) {
	var req user.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
	}

	resp, err := rpc.UserClient.Login(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// Register 用户注册
func (api *UserApi) Register(ctx *gin.Context) {
	var req user.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
	}

	resp, err := rpc.UserClient.Register(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// ChangePassword 修改密码
func (api *UserApi) ChangePassword(ctx *gin.Context) {
	var req user.ChangePasswordReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId

	resp, err := rpc.UserClient.ChangePassword(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// UpdateUserInfo 更新用户信息
func (api *UserApi) UpdateUserInfo(ctx *gin.Context) {
	var req user.UpdateUserInfoReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId

	resp, err := rpc.UserClient.UpdateUserInfo(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// GetUserInfo 获取用户信息
func (api *UserApi) GetUserInfo(ctx *gin.Context) {
	var req user.GetUserInfoReq

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)

	req.UserId = userId
	fmt.Println("userID", req.UserId)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	resp, err := rpc.UserClient.GetUserInfo(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}

// DeleteUser 删除用户
func (api *UserApi) DeleteUser(ctx *gin.Context) {
	var req user.DeleteUserReq

	// 设置 userId
	userId, err := domain.GetUserIdFromContext(ctx)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}
	req.UserId = userId

	resp, err := rpc.UserClient.DeleteUser(ctx, &req)
	if err != nil {
		domain.ErrorMsg(ctx, err.Error())
		return
	}

	domain.Success(ctx, resp)
}
