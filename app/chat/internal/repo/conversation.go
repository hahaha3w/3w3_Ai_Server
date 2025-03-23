package repo

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
)

func (r *ChatRepository) CreateConversation(ctx context.Context, conversation *domain.Conversation) error {
	log.Log().Debug(conversation)
	return r.db.Create(conversation).Error
}

func (r *ChatRepository) ListConversations(ctx context.Context, userID int, pageSize int, pageNumber int) ([]*domain.Conversation, error) {
	var conversations []*domain.Conversation

	if err := r.db.Model(&domain.Conversation{}).Where("user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	offset := (pageNumber - 1) * pageSize
	if err := r.db.Where("user_id = ?", userID).
		Offset(offset).Limit(pageSize).
		Find(&conversations).Error; err != nil {
		return nil, err
	}

	return conversations, nil
}
func (r *ChatRepository) UpdateConversationTime(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {

	if err := r.db.Model(&domain.Conversation{}).Where("conversation_id =?", conversation.ConversationID).
		Updates(map[string]interface{}{"updated_at": conversation.UpdatedAt}).Error; err != nil {
		return nil, err
	}
	return conversation, nil

}
func (r *ChatRepository) DeleteConversation(ctx context.Context, conversationID int, userID string) (err error) {

	err = r.db.Where("conversation_id = ? AND user_id = ?", conversationID, userID).Delete(&domain.Conversation{}).Error

	if err != nil {
		log.Log().Error(err)
		return err
	}
	return nil

}
