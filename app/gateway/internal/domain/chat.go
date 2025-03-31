package domain

type ConversationResp struct {
	ConversationId int32  `json:"conversation_id"`
	UserId         int32  `json:"user_id"`
	SessionTitle   string `json:"session_title"`
	Mode           string `json:"mode"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
type ListConversationsResp struct {
	Conversations []*ConversationResp `json:"conversations"`
	Total         int32               `json:"total"`
	HasMore       bool                `json:"has_more"`
}
type MessageResp struct {
	MessageId      int32  `json:"message_id"`
	ConversationId int32  `json:"conversation_id"`
	UserId         int32  `json:"user_id"`
	Content        string `json:"content"`
	SenderType     string `json:"sender_type"`
	SendTime       string `json:"send_time"`
}
type ListMessagesResp struct {
	Messages []*MessageResp `json:"messages"`
	Total    int32          `json:"total"`
	HasMore  bool           `json:"has_more"`
}
type DeleteMessageResp struct {
	Success bool `json:"success"`
}
