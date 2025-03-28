package domain

import (
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat"
	"time"
)

type Memoir struct {
	MemoirID  int       `gorm:"column:memoir_id;primaryKey;autoIncrement" json:"memoir_id"`        // 回忆录唯一标识
	UserID    int       `gorm:"column:user_id;not null;index:idx_user_start" json:"user_id"`       // 关联用户ID
	Title     string    `gorm:"column:title;not null" json:"title"`                                // 回忆录标题
	Type      string    `gorm:"column:type;not null;default:''" json:"type"`                       // 回忆录类型，有日记，周记，月记
	Style     string    `gorm:"column:style;not null;default:''" json:"style"`                     // 回忆录风格，有简约，文艺，记录
	StartDate time.Time `gorm:"column:start_date;not null;index:idx_user_start" json:"start_date"` // 开始日期
	EndDate   time.Time `gorm:"column:end_date;not null" json:"end_date"`                          // 结束日期
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`                  // 回忆录内容文本
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`     // 创建日期
}

// TableName 指定表名
func (Memoir) TableName() string {
	return "memoir"
}

// Message Reference by github.com/hahaha3w/3w3_Ai_Server/app/chat/internal/domain/model.go
type Message struct {
	MessageID      int             `gorm:"primaryKey;column:message_id" json:"message_id"`
	UserID         int             `gorm:"type:varchar(64);not null;column:user_id" json:"user_id"`
	ConversationID int             `gorm:"type:varchar(64);not null;column:conversation_id" json:"conversation_id"`
	Content        string          `gorm:"type:text;not null;column:content" json:"content"`
	SenderType     chat.SenderType `gorm:"type:varchar(20);not null;column:sender_type" json:"sender_type"`
	SendTime       time.Time       `gorm:"column:send_time" json:"send_time"`
}
