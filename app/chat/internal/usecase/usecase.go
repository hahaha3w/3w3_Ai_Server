package usecase

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain/enum"
	"io"
	"log"

	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/common/llm"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
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
	_, err = u.repo.StoreChatRecord(question, enum.User, "1")
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
	_, err = u.repo.StoreChatRecord(wholeContent, enum.Assistant, "1")
	if err != nil {
		return err
	}
	return nil
}
