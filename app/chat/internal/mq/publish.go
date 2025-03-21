package mq

import (
	"encoding/json"
	"github.com/cloudwego/kitex/server"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"github.com/nats-io/nats.go"
)

const (
	MessageRecordPublish = "message_record_publish"
)

func (c *ChatMQ) PublishMessage(msg *domain.Message) (err error) {
	jsonMsg, err := json.Marshal(*msg)
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}
	err = c.mq.Publish(MessageRecordPublish, jsonMsg)
	if err != nil {
		log.Log().Error(err.Error())
	}
	return nil
}

func (c *ChatMQ) Subscribe() (err error) {

	sub, err := c.mq.Subscribe(MessageRecordPublish, func(msg *nats.Msg) {
		message := &domain.Message{}
		err = json.Unmarshal(msg.Data, message)
		log.Log().Info(message)
		if err != nil {
			log.Log().Error(err.Error())
		}
		_, err = c.repo.StoreChatRecord(message)
		if err != nil {
			log.Log().Error(err.Error())
		}
	})
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}
	server.RegisterShutdownHook(func() {
		sub.Unsubscribe()
		c.mq.Close()
	})
	return nil
}
