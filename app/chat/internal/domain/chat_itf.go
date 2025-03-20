package domain

import "github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"

type Usecase interface {
	Chat(question string, answer chat.Echo_ChatServer) (err error)
}
type Repository interface {
	StoreChatRecord(content string, role string, containerID string) error
}
