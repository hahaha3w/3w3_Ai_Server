package domain

import (
	"context"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (err error)
	FindUserByEmail(ctx context.Context, email string) (user *User, err error)
}
type UserUsecase interface {
	RegisterUser(ctx context.Context, email, password string) (UserID int32, err error)
	LoginUser(ctx context.Context, email, password string) (UserID int32, err error)
}
