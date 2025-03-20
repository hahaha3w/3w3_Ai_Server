package domain

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain/enum"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

type Usecase interface {
	Chat(question string, answer chat.Echo_ChatServer) (err error)
}
type Repository interface {
	StoreChatRecord(content string, role enum.Role, conversationID string) (userID int, err error)
}
type MessageQueue[T any] interface {
	Publish(msg T) (err error)
}
