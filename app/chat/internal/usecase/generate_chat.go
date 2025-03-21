package usecase

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain/enum"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

type ChatUsecase struct {
	repo domain.Repository
	mq   domain.MessageQueue
}

func NewChatUsecase(repo domain.Repository, mq domain.MessageQueue) *ChatUsecase {
	return &ChatUsecase{
		repo: repo,
		mq:   mq,
	}
}
func (u *ChatUsecase) Chat(question string, answer chat.Echo_ChatServer) (err error) {
	userMessage := &domain.Message{
		ConversationID: "1",
		Content:        question,
		SenderType:     enum.User,
		SendTime:       time.Now(),
	}
	_, err = u.repo.StoreChatRecord(userMessage)
	if err != nil {
		return err
	}
	cm := llm.GetOpenAIChatModel()
	m := llm.CreateMessagesFromTemplate(question)
	sr, err := cm.Stream(context.Background(), m)
	if err != nil {
		log.Println(err)
		return
	}
	defer sr.Close()
	i := 0
	resp := &chat.ChatResp{}
	wholeContent := ""
	for {
		msg, err := sr.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		resp.Message = msg.Content
		wholeContent += msg.Content
		err = answer.Send(resp)
		if err != nil {
			log.Println(err)
			return err
		}
		i++
	}
	messageModel := domain.Message{
		Content:        wholeContent,
		SenderType:     enum.Assistant,
		SendTime:       time.Now(),
		ConversationID: "1",
	}
	err = u.mq.PublishMessage(&messageModel)
	if err != nil {
		return err
	}
	return nil
}
