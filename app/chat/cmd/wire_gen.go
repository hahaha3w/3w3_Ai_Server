// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/cache"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/delivery"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/mq"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/repo"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/usecase"
)

// Injectors from wire.go:

func wireApp() *delivery.ChatDelivery {
	context := core.NewContext()
	db := core.NewDB(context)
	conn := core.NewMQ()
	chatMQ := mq.NewChatMQ(conn)
	chatRepository := repo.NewChatRepository(db, chatMQ)
	client := core.NewRedis(context)
	chatRepoWithCache := cache.NewChatRepoWithCache(chatRepository, client, chatMQ)
	chatUsecase := usecase.NewChatUsecase(chatRepoWithCache, chatMQ)
	chatDelivery := delivery.NewChatDelivery(chatUsecase, chatMQ)
	return chatDelivery
}
