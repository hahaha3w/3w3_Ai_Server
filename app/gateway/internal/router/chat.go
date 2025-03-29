package router

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

func (g *Group) SetChatRouter() {
	chatGroup := g.Group("/chat")
	chatGroup.Use(middleware.JWT())
	{
		// 创建会话
		chatGroup.POST("/conversation", api.CreateConversation)
		// 获取会话列表
		chatGroup.GET("/conversation", api.ListConversations)
		// 删除会话
		chatGroup.DELETE("/conversation", api.DeleteConversation)
		// 发送消息
		chatGroup.GET("/message", api.SendMessage)
		// 获取消息列表
		chatGroup.GET("/message/list", api.ListMessages)
	}
}
