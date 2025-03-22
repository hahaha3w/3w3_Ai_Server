package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
)

var _ domain.ActivityUsecase = &ConcreteActivityUsecase{}

type ConcreteActivityUsecase struct {
	repo  domain.ActivityRepo
	cache *redis.Client
}

func NewConcreteActivityUsecase(repo domain.ActivityRepo, cache *redis.Client) *ConcreteActivityUsecase {
	return &ConcreteActivityUsecase{
		repo:  repo,
		cache: cache,
	}
}

func (c ConcreteActivityUsecase) CreateUserActivity(ctx context.Context, userId int64, relationId int64, ActivityType, description string) (activity *domain.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

func (c ConcreteActivityUsecase) GetUserActivity(ctx context.Context, userId int64) (activity *[]domain.Activity, err error) {
	//TODO implement me
	panic("implement me")
}
