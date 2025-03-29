package usecase

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"golang.org/x/net/context"
)

func (u *ChatUsecase) UpdateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {
	_, err := u.UpdateConversation(ctx, conversation)
	if err != nil {
		return nil, err
	}
	return conversation, nil

}
func (u *ChatUsecase) CreateConversation(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {

	err := u.repo.CreateConversation(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (u *ChatUsecase) ListConversations(ctx context.Context, userID int, pageSize int, pageNumber int) ([]*domain.Conversation, error) {
	return u.repo.ListConversations(ctx, userID, pageSize, pageNumber)
}

func (u *ChatUsecase) DeleteConversation(ctx context.Context, conversationID int, userID int) error {
	return u.repo.DeleteConversation(ctx, conversationID, userID)
}
