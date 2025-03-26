package api

import (
	"context"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

// CreateConversation 创建会话
func CreateConversation(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	var req chat.CreateConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		domain.Error(c, 400, err.Error())
		return
	}

	req.UserId = userID
	resp, err := rpc.ChatClient.CreateConversation(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	domain.Success(c, resp)
}

// ListConversations 获取会话列表
func ListConversations(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	var req chat.ListConversationsRequest
	pageSize := c.DefaultQuery("page_size", "10")
	pageNumber := c.DefaultQuery("page_number", "1")

	// 设置分页参数
	pageNum, err := strconv.Atoi(pageNumber)
	if err != nil {
		domain.Error(c, 400, "invalid page_number")
		return
	}
	pageS, err := strconv.Atoi(pageSize)
	if err != nil {
		domain.Error(c, 400, "invalid page_size")
		return
	}
	req.PageNumber = int32(pageNum)
	req.PageSize = int32(pageS)

	req.UserId = userID
	resp, err := rpc.ChatClient.ListConversations(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	domain.Success(c, resp)
}

// DeleteConversation 删除会话
func DeleteConversation(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	conversationID := c.Query("id")
	if conversationID == "" {
		domain.Error(c, 400, "conversation_id is required")
		return
	}

	convID, err := strconv.ParseInt(conversationID, 10, 32)
	if err != nil {
		domain.Error(c, 400, "invalid conversation_id")
		return
	}

	var req chat.DeleteConversationRequest
	req.ConversationId = int32(convID)
	req.UserId = userID

	resp, err := rpc.ChatClient.DeleteConversation(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	domain.Success(c, resp)
}

// SendMessage 发送消息
func SendMessage(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	var req chat.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		domain.Error(c, 400, err.Error())
		return
	}

	req.UserId = userID
	stream, err := rpc.ChatClient.SendMessage(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	c.Stream(func(w io.Writer) bool {
		resp, err := stream.Recv()
		if err != nil {
			return false
		}
		c.SSEvent("message", resp)
		return true
	})
}

// ListMessages 获取消息列表
func ListMessages(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	var req chat.ListMessagesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		domain.Error(c, 400, err.Error())
		return
	}

	req.UserId = userID
	resp, err := rpc.ChatClient.ListMessages(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	domain.Success(c, resp)
}
