package mq

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/repo"
	"github.com/nats-io/nats.go"
)

type ChatMQ struct {
	mq   *nats.Conn
	repo *repo.ChatRepository
}

var ProviderSet = wire.NewSet(wire.Bind(new(domain.MessageQueue), new(*ChatMQ)), NewChatMQ)

func NewChatMQ(mq *nats.Conn, repo *repo.ChatRepository) *ChatMQ {
	return &ChatMQ{
		mq:   mq,
		repo: repo,
	}
}
