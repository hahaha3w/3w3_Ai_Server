package usecase

import (
	"golang.org/x/net/context"
	"time"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
)

func NewConversationUsecase(repo domain.Repository) *ChatUsecase {
	return &ChatUsecase{repo: repo}
}

func (u *ChatUsecase) CreateConversation(ctx context.Context, userID int, sessionTitle string, mode string) (*domain.Conversation, error) {
	conversation := &domain.Conversation{
		UserID:       userID,
		SessionTitle: sessionTitle,
		Mode:         mode,
		CreatedAt:    time.Now(),
	}

	err := u.repo.CreateConversation(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (u *ChatUsecase) ListConversations(ctx context.Context, userID int, pageSize int, pageNumber int) ([]*domain.Conversation, error) {
	return u.repo.ListConversations(ctx, userID, pageSize, pageNumber)
}

func (u *ChatUsecase) DeleteConversation(ctx context.Context, conversationID int, userID string) error {
	return u.repo.DeleteConversation(ctx, conversationID, userID)
}
