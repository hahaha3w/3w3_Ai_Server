package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/repo"
)

type ChatRepoWithCache struct {
	redis *redis.Client
	repo  *repo.ChatRepository
	mq    domain.MessageQueue
}

var ProviderSet = wire.NewSet(wire.Bind(new(domain.Repository), new(*ChatRepoWithCache)), NewChatRepoWithCache)
var _ domain.Repository = &ChatRepoWithCache{}

func NewChatRepoWithCache(repo *repo.ChatRepository, cache *redis.Client, mq domain.MessageQueue) *ChatRepoWithCache {
	return &ChatRepoWithCache{
		repo:  repo,
		redis: cache,
		mq:    mq,
	}
}
