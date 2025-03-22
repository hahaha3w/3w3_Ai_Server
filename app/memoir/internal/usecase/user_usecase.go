package usecase

import (
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
)

var _ domain.Usecase = &ConcreteMemoirUsecase{}

type ConcreteMemoirUsecase struct {
	repo  domain.MemoirRepo
	cache *redis.Client
}

func NewConcreteMemoirUsecase(repo domain.MemoirRepo, cache *redis.Client) *ConcreteMemoirUsecase {
	return &ConcreteMemoirUsecase{
		repo:  repo,
		cache: cache,
	}
}
