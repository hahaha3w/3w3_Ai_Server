package usecase

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	"io"
	"log"
)

type ChatUsecase struct {
	repo domain.Repository
}

func NewChatUsecase(repo domain.Repository) *ChatUsecase {
	return &ChatUsecase{
		repo: repo,
	}
}
func (u *ChatUsecase) Chat(question string, answer chat.Echo_ChatServer) (err error) {
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
	for {
		msg, err := sr.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return err
		}
		resp.Message = msg.Content
		err = answer.Send(resp)
		if err != nil {
			log.Println(err)
			return err
		}
		i++
	}
	return nil
}
