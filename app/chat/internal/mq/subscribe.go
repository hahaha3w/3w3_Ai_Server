package mq

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/kitex/server"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"github.com/nats-io/nats.go"
)

func (c *ChatMQ) Subscribe() (err error) {
	msgSub, err := c.mq.Subscribe(MessageRecordPublish, func(msg *nats.Msg) {
		log.Log().Info("consume scribe message")
		message := &domain.Message{}
		err = json.Unmarshal(msg.Data, message)
		log.Log().Info(message)
		if err != nil {
			log.Log().Error(err.Error())
		}
		_, err = c.repo.StoreChatRecord(context.Background(), message)
		if err != nil {
			log.Log().Error(err.Error())
		}
	})
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}

	convSub, err := c.mq.Subscribe(ConversationRecordPublish, func(msg *nats.Msg) {
		log.Log().Info("consume scribe conversation")
		conv := &domain.Conversation{}
		err = json.Unmarshal(msg.Data, conv)
		log.Log().Info(conv)
		if err != nil {
			log.Log().Error(err.Error())
		}
		_, err = c.repo.UpdateConversationTime(context.Background(), conv)
		if err != nil {
			log.Log().Error(err.Error())
		}
	})
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}

	server.RegisterShutdownHook(func() {
		msgSub.Unsubscribe()
		convSub.Unsubscribe()
		c.mq.Close()
	})
	return nil
}
