package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/email"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/encrypt"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/regex"
	"time"

	"gorm.io/gorm"
)

var (
	errUserAlreadyExist = errors.New("user already exist")
	errPasswordNotMatch = errors.New("password does not match")
)
var _ domain.UserUsecase = &ConcreteUserUsecase{}

type ConcreteUserUsecase struct {
	repo  domain.UserRepo
	cache *redis.Client
}

func NewConcreteUserUsecase(repo domain.UserRepo, cache *redis.Client) *ConcreteUserUsecase {
	return &ConcreteUserUsecase{
		repo:  repo,
		cache: cache,
	}
}

func (u *ConcreteUserUsecase) SendCode(ctx context.Context, sendTo string) (err error) {
	// 校验邮箱
	if regex.IsEmailInvalid(sendTo) {
		log.Log().Error(err)
		return errors.New("invalid email")
	}

	// 是否 10 min 内重复发送验证码
	code := u.cache.Get(ctx, "email_code:"+sendTo).Val()
	if code != "" {
		log.Log().Error(err)
		return errors.New("code already sent, please not resend within 10 min")
	}

	// 随机生成 6 位数字验证码
	code = email.RandomNumbers(6)

	// 将验证码存入 redis，10 min 后过期
	log.Log().Info("用户 " + sendTo + " 验证码为 " + code)
	err = u.cache.Set(ctx, "email_code:"+sendTo, code, 10*60*time.Second).Err()
	if err != nil {
		log.Log().Error(err)
		return errors.New("internal error: " + err.Error())
	}

	// 使用发送验证码
	go func() {
		_ = email.SendVerificationCode(sendTo, code)
	}()

	return nil
}

func (u *ConcreteUserUsecase) RegisterUser(ctx context.Context, email, password string) (userID int32, err error) {
	var (
		userExist *domain.User
		user      *domain.User
	)

	// check if user email already exist
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
