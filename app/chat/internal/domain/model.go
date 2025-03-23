package domain

import (
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	"time"
)

// Conversation 对话容器表
type Conversation struct {
	ConversationID int       `gorm:"primaryKey;column:conversation_id" json:"conversation_id"`
	UserID         int       `gorm:"type:varchar(64);not null;column:user_id" json:"user_id"`
	SessionTitle   string    `gorm:"type:varchar(255);default:'';column:session_title" json:"session_title"`
	Mode           string    `gorm:"type:varchar(20);default:'';column:mode" json:"mode"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt      time.Time `gorm:"column:create_time" json:"created_at"`
}

// Message 消息内容表
type Message struct {
	MessageID      int             `gorm:"primaryKey;column:message_id" json:"message_id"`
	UserID         int             `gorm:"type:varchar(64);not null;column:user_id" json:"user_id"`
	ConversationID int             `gorm:"type:varchar(64);not null;column:conversation_id" json:"conversation_id"`
	Content        string          `gorm:"type:text;not null;column:content" json:"content"`
	SenderType     chat.SenderType `gorm:"type:varchar(20);not null;column:sender_type" json:"sender_type"`
	SendTime       time.Time       `gorm:"column:send_time" json:"send_time"`
}
