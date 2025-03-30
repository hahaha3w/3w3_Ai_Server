package api

import (
	"context"
	"fmt"
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
	vo := domain.ConversationResp{
		ConversationId: resp.Conversation.ConversationId,
		UserId:         resp.Conversation.UserId,
		SessionTitle:   resp.Conversation.SessionTitle,
		Mode:           resp.Conversation.Mode,
		CreateTime:     resp.Conversation.CreatedAt,
	}
	domain.Success(c, vo)
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
	conversations := make([]*domain.ConversationResp, 0)
	for _, v := range resp.Conversations {
		conversation := &domain.ConversationResp{
			ConversationId: v.ConversationId,
			UserId:         v.UserId,
			SessionTitle:   v.SessionTitle,
			Mode:           v.Mode,
			CreateTime:     v.CreatedAt,
		}
		conversations = append(conversations, conversation)
	}

	vo := domain.ListConversationsResp{
		Conversations: conversations,
		Total:         resp.Total,
	}

	domain.Success(c, vo)
}

// DeleteConversation 删除会话
func DeleteConversation(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	conversationID := c.Query("conversation_id")
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
	vo := domain.DeleteMessageResp{
		Success: resp.Success,
	}

	domain.Success(c, vo)
}

// SendMessage 发送消息
func SendMessage(c *gin.Context) {
	userID, err := domain.GetUserIdFromContext(c)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}

	req := &chat.SendMessageRequest{}
	//  ConversationId int32  `protobuf:"varint,1,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	//    UserId         int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	//    Content        string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	//}
	//
	//发送消息请求
	conversationID := c.Query("conversation_id")
	if conversationID == "" {
		domain.Error(c, 400, "conversation_id is required")
		return
	}
	convID, err := strconv.Atoi(conversationID)
	if err != nil {
		domain.Error(c, 400, "invalid conversation_id")
		return
	}
	content := c.Query("content")
	if content == "" {
		domain.Error(c, 400, "content is required")
		return
	}
	req.Content = content
	req.ConversationId = int32(convID)
	req.UserId = userID
	stream, err := rpc.ChatClient.SendMessage(context.Background(), req)
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
		if err == io.EOF {
			return true
		}
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
		fmt.Println(err)
		domain.ErrorMsg(c, err.Error())
		return
	}
	pageSize := c.DefaultQuery("page_size", "10")
	pageNumber := c.DefaultQuery("page_number", "1")
	convID, err := strconv.Atoi(c.Query("conversation_id"))
	fmt.Println(convID, pageSize, pageNumber)
	if err != nil {
		domain.Error(c, 400, "invalid conv_id")
		return
	}
	pageS, err := strconv.Atoi(pageSize)
	if err != nil {
		domain.Error(c, 400, "invalid page_size")
		return
	}
	pageNum, err := strconv.Atoi(pageNumber)
	if err != nil {
		domain.Error(c, 400, "invalid page_number")
		return
	}
	req := chat.ListMessagesRequest{
		UserId:         userID,
		PageSize:       int32(pageS),
		PageNumber:     int32(pageNum),
		ConversationId: int32(convID),
	}
	resp, err := rpc.ChatClient.ListMessages(context.Background(), &req)
	if err != nil {
		domain.ErrorMsg(c, err.Error())
		return
	}
	messages := make([]*domain.MessageResp, 0)
	for _, m := range resp.Messages {
		message := &domain.MessageResp{
			MessageId:      m.MessageId,
			ConversationId: m.ConversationId,
			UserId:         m.UserId,
			Content:        m.Content,
			SenderType:     m.SenderType,
			SendTime:       m.SendTime,
		}
		messages = append(messages, message)
	}
	vo := domain.ListMessagesResp{
		Messages: messages,
		Total:    resp.Total,
	}
	domain.Success(c, vo)
}
