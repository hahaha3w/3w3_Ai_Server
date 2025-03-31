package repo

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"gorm.io/gorm"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
)

func (r *ChatRepository) StoreMessageRecord(ctx context.Context, m *domain.Message) (MessageID int, err error) {
	err = r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Create(m).Error; err != nil {
			log.Log().Error(err.Error())
			return err
		}
		err = r.mq.PublishConversation(m.ConversationID, m.SendTime)
		if err != nil {
			log.Log().Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Log().Error(err.Error())
		return 0, err
	}
	return m.MessageID, nil
}
func (r *ChatRepository) ListMessages(ctx context.Context, conversationID, userID, start, end int) ([]*domain.Message, error) {
	var messages []*domain.Message

	if err := r.db.Where("conversation_id = ? AND user_id = ?", conversationID, userID).
		Order("send_time DESC").
		Offset(start).
		Limit(end - start + 1).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}
