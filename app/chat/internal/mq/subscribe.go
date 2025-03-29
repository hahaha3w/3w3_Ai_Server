package mq

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/kitex/server"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"github.com/nats-io/nats.go"
)

func (c *ChatMQ) Subscribe(RepoCallback func(ctx context.Context, conv *domain.Conversation) (*domain.Conversation, error)) (err error) {

	convSub, err := c.mq.Subscribe(ConversationRecordPublish, func(msg *nats.Msg) {
		log.Log().Info("consume scribe conversation")
		conv := &domain.Conversation{}
		err = json.Unmarshal(msg.Data, conv)
		log.Log().Info(conv)
		if err != nil {
			log.Log().Error(err.Error())
		}
		_, err = RepoCallback(context.Background(), conv)
		if err != nil {
			log.Log().Error(err.Error())
		}
	})
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}

	server.RegisterShutdownHook(func() {
		convSub.Unsubscribe()
		c.mq.Close()
	})
	return nil
}
