package domain

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (err error)
	FindUserByEmail(ctx context.Context, email string) (user *User, err error)
	FindUserByID(ctx context.Context, id int64) (user *User, err error)
	UpdateUser(ctx context.Context, user *User) (err error)
}

type UserUsecase interface {
	SendCode(ctx context.Context, email string) (err error)
	RegisterUser(ctx context.Context, email, code, password string) (resp *user.RegisterResp, err error)
	LoginUser(ctx context.Context, email, password string) (resp *user.LoginResp, err error)
	GetUser(ctx context.Context, id int64) (resp *user.GetUserInfoResp, err error)
	UpdatePassword(ctx context.Context, id int32, oldPassword, newPassword string) (err error)
	UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (err error)
}
