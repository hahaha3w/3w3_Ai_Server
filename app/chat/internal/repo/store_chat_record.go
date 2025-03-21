package repo

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}
func (r *ChatRepo) StoreChatRecord(m *domain.Message) (userID int, err error) {
	
	err = r.db.Create(m).Error
	if err != nil {
		return 0, err
	}
	return m.MessageID, nil
}
