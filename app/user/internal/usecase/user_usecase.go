package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/count"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/email"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/encrypt"
	"github.com/hahaha3w/3w3_Ai_Server/user/pkg/jwt"
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

func (u *ConcreteUserUsecase) RegisterUser(ctx context.Context, email, code, password string) (resp *user.RegisterResp, err error) {
	var (
		userExist *domain.User
		userModel *domain.User
	)

	// 校验邮箱和密码格式
	err = regex.VerifyUser(email, password)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	// 检查用户邮箱是否已存在
	userExist, err = u.repo.FindUserByEmail(ctx, email)
	if userExist.UserID != 0 {
		log.Log().Info(errUserAlreadyExist)
		return nil, errors.New("user already exist")
	}

	// 校验验证码
	OriginCode := u.cache.Get(ctx, "email_code:"+email).Val()
	if code == "" || OriginCode != code {
		log.Log().Error(err)
		return nil, errors.New("verification code is incorrect, please check again")
	}

	// 哈希密码
	passwordHash, err := encrypt.HashPassword(password)
	if err != nil {
		log.Log().Error(err)
		return nil, errors.New("internal error: " + err.Error())
	}

	// 构建新用户
	userModel = &domain.User{
		Username:     "默认用户",
		Bio:          "该用户还没有没有填写个人简介",
		AvatarUrl:    "https://avatars.githubusercontent.com/u/204012462?s=48&v=4",
		Theme:        "light",
		RegisterDate: time.Now(),
		Email:        email,
		Password:     passwordHash,
	}

	// 创建用户
	err = u.repo.CreateUser(ctx, userModel)
	if err != nil {
		log.Log().Error(err)
		return nil, errors.New("internal error: " + err.Error())
	}

	// 使用哈希封装用户次数统计
	err = count.InitUserCount(userModel.UserID, u.cache)
	if err != nil {
		log.Log().Error(err)
		return nil, errors.New("internal error: " + err.Error())
	}

	// 返回用户信息
	return &user.RegisterResp{
		UserId:      userModel.UserID,
		Username:    userModel.Username,
		Email:       userModel.Email,
		Bio:         userModel.Bio,
		Avatar:      userModel.AvatarUrl,
		Theme:       userModel.Theme,
		ChatCount:   0,
		MemoirCount: 0,
		UseDay:      1,
		Token:       jwt.GenerateJWT(userModel.UserID),
	}, nil
}

func (u *ConcreteUserUsecase) LoginUser(ctx context.Context, email, password string) (resp *user.LoginResp, err error) {
	// 校验邮箱和密码格式
	err = regex.VerifyUser(email, password)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	// 根据邮箱查找用户
	userModel, err := u.repo.FindUserByEmail(ctx, email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Log().Info(err)
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("internal error: %w", err)
	}

	// 校验密码
	if !encrypt.ComparePasswords(userModel.Password, password) {
		log.Log().Info(errPasswordNotMatch)
		return nil, fmt.Errorf("password not match")
	}

	// 更新用户计数统计
	userCount := &domain.UpdateUserCount{
		UserId:          userModel.UserID,
		ChatCountIncr:   0,
		MemoirCountIncr: 0,
		UseDaysIncr:     1,
	}
	updateUserCount, err := count.UpdateUserCount(userCount, u.cache)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("internal error: %w", err)
	}

	// 返回用户信息
	return &user.LoginResp{
		UserId:      userModel.UserID,
		Username:    userModel.Username,
		Email:       userModel.Email,
		Bio:         userModel.Bio,
		Avatar:      userModel.AvatarUrl,
		Theme:       userModel.Theme,
		ChatCount:   int32(updateUserCount.ChatCount),
		MemoirCount: int32(updateUserCount.MemoirCount),
		UseDay:      int32(updateUserCount.UseDays),
		Token:       jwt.GenerateJWT(userModel.UserID),
	}, nil
}

func (u *ConcreteUserUsecase) GetUser(ctx context.Context, id int64) (resp *user.GetUserInfoResp, err error) {
	userModel, err := u.repo.FindUserByID(ctx, id)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}
	if userModel.UserID == 0 {
		return nil, errors.New("user not found")
	}

	userCount, err := count.GetUserCount(userModel.UserID, u.cache)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	return &user.GetUserInfoResp{
		UserId:      userModel.UserID,
		Username:    userModel.Username,
		Email:       userModel.Email,
		Bio:         userModel.Bio,
		Avatar:      userModel.AvatarUrl,
		Theme:       userModel.Theme,
		ChatCount:   int32(userCount.ChatCount),
		MemoirCount: int32(userCount.MemoirCount),
		UseDay:      int32(userCount.UseDays),
	}, nil
}

func (u *ConcreteUserUsecase) UpdatePassword(ctx context.Context, id int32, oldPassword, newPassword string) (err error) {
	// 校验密码格式
	if regex.IsPasswordInvalid(newPassword) {
		log.Log().Error(err)
		return errors.New("password invalid")
	}

	// 根据id查找用户
	userModel, err := u.repo.FindUserByID(ctx, int64(id))
	if err != nil {
		log.Log().Error(err)
		return err
	}
	if userModel.UserID == 0 {
		return errors.New("user not found")
	}

	// 校验密码
	if !encrypt.ComparePasswords(userModel.Password, oldPassword) {
		log.Log().Info(errPasswordNotMatch)
		return fmt.Errorf("password not match")
	}

	// 更新密码
	password, err := encrypt.HashPassword(newPassword)
	if err != nil {
		log.Log().Error(err)
		return err
	}
	userModel.Password = password
	if err := u.repo.UpdateUser(ctx, userModel); err != nil {
		log.Log().Error(err)
		return err
	}

	return nil
}

// UpdateUserInfo 更新用户信息
func (u *ConcreteUserUsecase) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoReq) (err error) {
	// 根据id查找用户
	userModel, err := u.repo.FindUserByID(ctx, int64(req.UserId))
	if err != nil {
		log.Log().Error(err)
		return err
	}

	if userModel.UserID == 0 {
		return errors.New("user not found")
	}

	if req.Username != "" {
		userModel.Username = req.Username
	}
	if req.Bio != "" {
		userModel.Bio = req.Bio
	}
	if req.Avatar != "" {
		userModel.AvatarUrl = req.Avatar
	}
	if req.Theme != "" {
		userModel.Theme = req.Theme
	}

	// 更新用户信息
	if err := u.repo.UpdateUser(ctx, userModel); err != nil {
		log.Log().Error(err)
		return err
	}

	return nil
}

func (u *ConcreteUserUsecase) DeleteUser(ctx context.Context, id int64) (err error) {
	if err := u.repo.DeleteUser(ctx, id); err != nil {
		log.Log().Error(err)
		return err
	}

	err = count.DeleteUserCount(int32(id), u.cache)
	if err != nil {
		log.Log().Error(err)
		return err
	}

	return nil
}
