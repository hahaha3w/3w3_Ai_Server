package cache

import (
	"context"
	"fmt"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
)

func (c *ChatRepoWithCache) UpdateConversationTime(ctx context.Context, conversation *domain.Conversation) (*domain.Conversation, error) {
	return c.repo.UpdateConversationTime(ctx, conversation)

}

// CreateConversation 将新的会话添加到用户的会话列表中
func (c *ChatRepoWithCache) CreateConversation(ctx context.Context, conversation *domain.Conversation) error {
	// 生成用户会话列表的key
	err := c.repo.CreateConversation(ctx, conversation)
	if err != nil {
		log.Log().Error(err)
		return err
	}

	return err
}

// ListConversations 获取用户的会话列表
func (c *ChatRepoWithCache) ListConversations(ctx context.Context, userID int, pageSize, pageNumber int) ([]*domain.Conversation, error) {
	// 生成用户会话列表的key

	conversations, err := c.repo.ListConversations(ctx, userID, pageSize, pageNumber)
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("failed to get conversations: %w", err)
	}

	return conversations, nil
}

// DeleteConversation 从用户的会话列表中删除指定的会话
func (c *ChatRepoWithCache) DeleteConversation(ctx context.Context, conversationID int, userID int) error {
	// 生成用户会话列表的key

	err := c.repo.DeleteConversation(ctx, conversationID, userID)
	if err != nil {
		log.Log().Error(err)
		return fmt.Errorf("failed to delete conversation: %w", err)
	}

	return nil
}
