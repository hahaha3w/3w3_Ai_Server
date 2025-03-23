package mq

import (

	"encoding/json"
	"time"


	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"

)

const (
	MessageRecordPublish      = "message_record_publish"
	ConversationRecordPublish = "conversation_record_publish"
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
func (c *ChatMQ) PublishConversation(conversationID int, newUpdateTime time.Time) (err error) {
	conv := &domain.Conversation{
		ConversationID: conversationID,
		UpdatedAt:      newUpdateTime,
	}
	jsonConv, err := json.Marshal(conv)
	if err != nil {
		log.Log().Error(err.Error())
		return err
	}
	err = c.mq.Publish(ConversationRecordPublish, jsonConv)
	if err != nil {
		log.Log().Error(err.Error())
	}
	return nil
}
