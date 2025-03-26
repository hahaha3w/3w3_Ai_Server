package delivery

import (
	"context"

	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
)

func (d *ChatDelivery) CreateConversation(ctx context.Context, req *chat.CreateConversationRequest) (res *chat.CreateConversationResponse, err error) {
	conversation, err := d.usecase.CreateConversation(ctx, int(req.UserId), req.SessionTitle, req.Mode)
	if err != nil {
		return nil, err
	}

	return &chat.CreateConversationResponse{
		Conversation: &chat.Conversation{
			ConversationId: int32(conversation.ConversationID),
			UserId:         int32(conversation.UserID),
			SessionTitle:   conversation.SessionTitle,
			Mode:           conversation.Mode,
			CreateTime:     conversation.CreatedAt.Format("2006-01-02T15:04:05Z"),
		},
	}, nil
}

func (d *ChatDelivery) ListConversations(ctx context.Context, req *chat.ListConversationsRequest) (res *chat.ListConversationsResponse, err error) {
	conversations, err := d.usecase.ListConversations(ctx, int(req.UserId), int(req.PageSize), int(req.PageNumber))
	if err != nil {
		return nil, err
	}

	var protoConversations []*chat.Conversation
	for _, conv := range conversations {
		protoConversations = append(protoConversations, &chat.Conversation{
			ConversationId: int32(conv.ConversationID),
			UserId:         int32(conv.UserID),
			SessionTitle:   conv.SessionTitle,
			Mode:           conv.Mode,
			CreateTime:     conv.CreatedAt.Format("2006-01-02T15:04:05Z"),
		})
	}

	return &chat.ListConversationsResponse{
		Conversations: protoConversations,
		Total:         int32(len(conversations)),
	}, nil
}

func (d *ChatDelivery) DeleteConversation(ctx context.Context, req *chat.DeleteConversationRequest) (res *chat.DeleteConversationResponse, err error) {
	err = d.usecase.DeleteConversation(ctx, int(req.ConversationId), int(req.UserId))
	if err != nil {
		return &chat.DeleteConversationResponse{
			Success: false,
		}, err
	}

	return &chat.DeleteConversationResponse{
		Success: true,
	}, nil
}
