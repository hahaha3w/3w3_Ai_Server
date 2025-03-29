package delivery

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/chat/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

func (d *ChatDelivery) InitSubscribe() {
	err := d.mq.Subscribe(d.usecase.CreateConversation)
	if err != nil {
		log.Log().Error(err)
	}
}

func (d *ChatDelivery) SendMessage(req *chat.SendMessageRequest, stream chat.ChatService_SendMessageServer) (err error) {
	err = d.usecase.SendMessage(context.Background(), req, stream)
	if err != nil {
		log.Log().Error(err)
	}
	return nil
}

func (d *ChatDelivery) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (res *chat.ListMessagesResponse, err error) {
	// 调用usecase层获取消息列表
	messages, err := d.usecase.ListMessages(ctx, req)
	if err != nil {
		log.Log().Error(err)
		return nil, err
	}

	// 将domain.Message转换为proto.Message
	var protoMessages []*chat.Message
	for _, msg := range messages {
		protoMessages = append(protoMessages, &chat.Message{
			MessageId:      int32(msg.MessageID),
			ConversationId: int32(msg.ConversationID),
			Content:        msg.Content,
			SenderType:     string(msg.SenderType),
			SendTime:       msg.SendTime.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &chat.ListMessagesResponse{
		Messages: protoMessages,
		Total:    int32(len(messages)),
	}, nil
}
