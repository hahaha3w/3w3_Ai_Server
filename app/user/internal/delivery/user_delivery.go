package delivery

import (
	"context"
	"errors"
	"fmt"
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

// Register implements the UserServiceImpl interface.
func (d *UserDelivery) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	var (
		uid int32
	)

	if req.Password != req.ConfirmPassword {
		log.Log().Error(errPasswordNotMatch)
		return nil, fmt.Errorf("delivery:%w", errPasswordNotMatch)
	}
	uid, err = d.userUsecase.RegisterUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("delivery:%w", err)
	}
	return &user.RegisterResp{
		UserId: uid,
	}, nil
}

// Login implements the UserServiceImpl interface.
func (d *UserDelivery) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	var (
		uid int32
	)

	uid, err = d.userUsecase.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("delivery:%w", err)
	}

	return &user.LoginResp{
		UserId: uid,
	}, err
}
func (d *UserDelivery) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (res *user.UpdateUserInfoResp, err error) {
	return nil, err
}
func (d *UserDelivery) GetUserInfo(ctx context.Context, req *user.GetUserInfoReq) (res *user.GetUserInfoResp, err error) {
	return nil, err
}

func (d *UserDelivery) DeleteUser(ctx context.Context, req *user.DeleteUserReq) (res *user.DeleteUserResp, err error) {
	return nil, err
}
