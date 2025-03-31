package repo

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewChatRepository)
var _ domain.Repository = &ChatRepository{}

type ChatRepository struct {
	db *gorm.DB
	mq domain.MessageQueue
}

func NewChatRepository(db *gorm.DB, mq domain.MessageQueue) *ChatRepository {
	return &ChatRepository{
		db: db,
		mq: mq,
	}
}
