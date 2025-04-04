package usecase

import (
	"context"
	"fmt"
	"io"

	"time"

	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/llm"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

func NewChatUsecase(repo domain.Repository, mq domain.MessageQueue) *ChatUsecase {
	return &ChatUsecase{
		repo: repo,
		mq:   mq,
	}
}
func (u *ChatUsecase) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) ([]*domain.Message, error) {

	// 从Redis缓存中获取消息列表
	result, err := u.repo.ListMessages(ctx, int(req.ConversationId), int(req.UserId), int(req.PageNumber), int(req.PageSize))
	if err != nil {
		log.Log().Error(err)
		return nil, fmt.Errorf("failed to get messages from cache: %v", err)
	}

	return result, nil
}
func (u *ChatUsecase) SendMessage(ctx context.Context, req *chat.SendMessageRequest, stream chat.ChatService_SendMessageServer) (err error) {
	userMessage := &domain.Message{
		UserID:         int(req.UserId),
		ConversationID: int(req.GetConversationId()),
		Content:        req.GetContent(),
		SenderType:     chat.SenderType_SENDER_USER,
		SendTime:       time.Now(),
	}
	_, err = u.repo.StoreMessageRecord(context.Background(), userMessage)
	if err != nil {
		log.Log().Error(err)
		return err
	}

	sr, err := llm.Chat(ctx, req.GetContent(), req.Mode)
	if err != nil {
		log.Log().Error(err)
		return
	}
	defer sr.Close()
	i := 0
	resp := &chat.SendMessageResponse{
		Message: &chat.Message{},
	}
	wholeContent := ""
	for {
		msg, err := sr.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Log().Error(err)
			return err
		}
		resp.Message.Content = msg.Content
		wholeContent += msg.Content
		err = stream.Send(resp)
		if err != nil {
			log.Log().Error(err)
			return err
		}
		i++
	}

	messageModel := domain.Message{
		UserID:         int(req.UserId),
		Content:        wholeContent,
		SenderType:     chat.SenderType_SENDER_AI,
		SendTime:       time.Now(),
		ConversationID: userMessage.ConversationID,
	}
	_, err = u.repo.StoreMessageRecord(ctx, &messageModel)
	if err != nil {
		log.Log().Error(err)
		return err

	}
	return nil
}
