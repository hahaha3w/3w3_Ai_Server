package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/encrypt"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/log"

	"gorm.io/gorm"
)

var (
	errUserAlreadyExist = errors.New("user already exist")
	errPasswordNotMatch = errors.New("password does not match")
)
var _ domain.UserUsecase = &ConcreteUserUsecase{}

type ConcreteUserUsecase struct {
	repo domain.UserRepo
}

func NewConcreteUserUsecase(repo domain.UserRepo) *ConcreteUserUsecase {
	return &ConcreteUserUsecase{
		repo: repo,
	}

}

func (u *ConcreteUserUsecase) RegisterUser(ctx context.Context, email, password string) (userID int32, err error) {
	var (
		userExist *domain.User
		user      *domain.User
	)

	userExist, err = u.repo.FindUserByEmail(ctx, email)
	if userExist != nil {
		log.Log().Info(errUserAlreadyExist)
		return 0, fmt.Errorf("usecase.reg:%w", errUserAlreadyExist)
	}
	passwordHash, err := encrypt.HashPassword(password)
	if err != nil {
		log.Log().Error(err)
		return 0, fmt.Errorf("usecase.reg.encrypt:%w", err)
	}
	user = &domain.User{
		Email:          email,
		PasswordHashed: passwordHash,
	}
	err = u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Log().Error(err)
		return 0, fmt.Errorf("usecase.reg.create:%w", err)
	}
	return user.UserID, nil
}
func (u *ConcreteUserUsecase) LoginUser(ctx context.Context, email, password string) (UserID int32, err error) {
	user, err := u.repo.FindUserByEmail(ctx, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log().Info(err)
		return 0, fmt.Errorf("usecase.login:%w", err)
	}
	if err != nil {
		log.Log().Error(err)
		return 0, fmt.Errorf("usecase.login:%w", err)
	}
	if encrypt.ComparePasswords(user.PasswordHashed, password) {
		log.Log().Info(errPasswordNotMatch)
		return 0, fmt.Errorf("usecase.login:%w", errPasswordNotMatch)
	}
	return user.UserID, nil
}
