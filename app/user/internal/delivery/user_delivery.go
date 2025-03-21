package delivery

import (
	"context"
	"errors"
	"fmt"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/log"
)

var (
	errPasswordNotMatch = errors.New("password does not match")
)

type UserDelivery struct {
	userUsecase domain.UserUsecase
}

func NewUserDelivery(userUsecase domain.UserUsecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: userUsecase,
	}
}

// SendCode implements the UserServiceImpl interface.
func (d *UserDelivery) SendCode(ctx context.Context, req *user.SendCodeReq) (resp *user.SendCodeResp, err error) {
	err = d.userUsecase.SendCode(ctx, req.Email)
	if err != nil {
		return &user.SendCodeResp{
			Success: false,
			Message: err.Error(),
		}, err
	}

	return &user.SendCodeResp{
		Success: true,
		Message: "send code success",
	}, nil
}

// Register implements the UserServiceImpl interface.
func (d *UserDelivery) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 重复校验密码
	if req.Password != req.ConfirmPassword {
		log.Log().Error(errPasswordNotMatch)
		return nil, fmt.Errorf("password does not match")
	}

	// 执行注册用户逻辑
	registerUser, err := d.userUsecase.RegisterUser(ctx, req.Email, req.Code, req.Password)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("register:%w", err)
	}

	return registerUser, nil
}

// Login implements the UserServiceImpl interface.
func (d *UserDelivery) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	loginUser, err := d.userUsecase.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("login:%w", err)
	}

	return loginUser, err
}

// UpdateUserInfo implements the UserServiceImpl interface.
func (d *UserDelivery) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (res *user.UpdateUserInfoResp, err error) {
	err = d.userUsecase.UpdateUserInfo(ctx, req)
	if err != nil {
		log.Log().Error(err)
		return &user.UpdateUserInfoResp{
			Success: false,
			Message: err.Error(),
		}, fmt.Errorf("update user info:%w", err)
	}

	return &user.UpdateUserInfoResp{
		Success: true,
		Message: "update user info success",
	}, nil
}

// GetUserInfo implements the UserServiceImpl interface.
func (d *UserDelivery) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (res *user.GetUserInfoResp, err error) {
	resp, err := d.userUsecase.GetUser(ctx, int64(req.UserId))
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("get user info:%w", err)
	}
	return resp, nil
}

// DeleteUser implements the UserServiceImpl interface.
func (d *UserDelivery) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (res *user.DeleteUserResp, err error) {
	err = d.userUsecase.DeleteUser(ctx, int64(req.UserId))
	if err != nil {
		log.Log().Error(err)
		return &user.DeleteUserResp{
			Success: false,
			Message: err.Error(),
		}, fmt.Errorf("delete user:%w", err)
	}

	return &user.DeleteUserResp{
		Success: true,
		Message: "delete user success",
	}, nil
}

// ChangePassword implements the UserServiceImpl interface.
func (d *UserDelivery) ChangePassword(ctx context.Context, req *user.ChangePasswordReq) (res *user.ChangePasswordResp, err error) {
	err = d.userUsecase.UpdatePassword(ctx, req.UserId, req.OldPassword, req.NewPassword)
	if err != nil {
		log.Log().Error(err)
		return &user.ChangePasswordResp{
			Success: false,
			Message: err.Error(),
		}, fmt.Errorf("change password:%w", err)
	}

	return &user.ChangePasswordResp{
		Success: true,
		Message: "change password success",
	}, nil
}
