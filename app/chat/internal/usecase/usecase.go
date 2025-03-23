package usecase

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
)

type ChatUsecase struct {
	repo domain.Repository
	mq   domain.MessageQueue
}

var ProviderSet = wire.NewSet(wire.Bind(new(domain.Usecase), new(*ChatUsecase)), NewChatUsecase)
var _ domain.Usecase = &ChatUsecase{}
