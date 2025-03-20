package repo

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain/enum"
	"gorm.io/gorm"
	"time"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}
func (r *ChatRepo) StoreChatRecord(content string, role enum.Role, conversationID string) (userID int, err error) {
	m := &domain.Message{
		Content:        content,
		SenderType:     role,
		ConversationID: conversationID,
		SendTime:       time.Now(),
	}
	err = r.db.Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.MessageID, nil
}
