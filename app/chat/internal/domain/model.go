package domain

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/domain/enum"
	"time"
)

// Conversation 对话容器表
type Conversation struct {
	ConversationID int       `gorm:"primaryKey;column:conversation_id" json:"conversation_id"`
	UserID         string    `gorm:"type:varchar(64);not null;column:user_id" json:"user_id"`
	SessionTitle   string    `gorm:"type:varchar(255);default:'';column:session_title" json:"session_title"`
	Mode           string    `gorm:"type:varchar(20);default:'';column:mode" json:"mode"`
	CreateAt       time.Time `gorm:"column:create_time" json:"create_at"`
}

// Message 消息内容表
type Message struct {
	MessageID      int       `gorm:"primaryKey;column:message_id" json:"message_id"`
	ConversationID string    `gorm:"type:varchar(64);not null;column:conversation_id" json:"conversation_id"`
	Content        string    `gorm:"type:text;not null;column:content" json:"content"`
	SenderType     enum.Role `gorm:"type:varchar(20);not null;column:sender_type" json:"sender_type"`
	SendTime       time.Time `gorm:"column:send_time" json:"send_time"`
}
