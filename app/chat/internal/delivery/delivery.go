package delivery

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

var _ chat.ChatService = &ChatDelivery{}
var ProviderSet = wire.NewSet(NewChatDelivery)

type ChatDelivery struct {
	usecase domain.Usecase
	mq      domain.MessageQueue
}

func NewChatDelivery(usecase domain.Usecase, mq domain.MessageQueue) *ChatDelivery {
	return &ChatDelivery{
		usecase: usecase,
		mq:      mq,
	}
}
