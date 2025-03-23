package domain

import (
	"context"
	"time"

	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

type Usecase interface {
	ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (res []*Message, err error)
	SendMessage(ctx context.Context, req *chat.SendMessageRequest, stream chat.ChatService_SendMessageServer) (err error)
	CreateConversation(ctx context.Context, userID int, sessionTitle string, mode string) (*Conversation, error)
	ListConversations(ctx context.Context, userID int, pageSize int, pageNumber int) ([]*Conversation, error)
	DeleteConversation(ctx context.Context, conversationID int, userID string) error
}
type Repository interface {
	StoreMessageRecord(ctx context.Context, m *Message) (userID int, err error)
	ListMessages(ctx context.Context, conversationID int, start int, end int) ([]*Message, error)
	CreateConversation(ctx context.Context, conversation *Conversation) error
	ListConversations(ctx context.Context, userID int, pageSize int, pageNumber int) ([]*Conversation, error)
	DeleteConversation(ctx context.Context, conversationID int, userID string) error
}

type MessageQueue interface {
	PublishMessage(msg *Message) (err error)
	PublishConversation(conversationID int, newUpdateTime time.Time) (err error)
	Subscribe() (err error)
}
