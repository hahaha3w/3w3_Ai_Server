package domain

import (
	"time"
)

type Memoir struct {
	MemoirID  int       `gorm:"column:memoir_id;primaryKey;autoIncrement" json:"memoir_id"`        // 回忆录唯一标识
	UserID    int       `gorm:"column:user_id;not null;index:idx_user_start" json:"user_id"`       // 关联用户ID
	Title     string    `gorm:"column:title;not null" json:"title"`                                // 回忆录标题
	Type      string    `gorm:"column:type;not null;default:''" json:"type"`                       // 回忆录类型，有日记，周记，月记
	StartDate time.Time `gorm:"column:start_date;not null;index:idx_user_start" json:"start_date"` // 开始日期
	EndDate   time.Time `gorm:"column:end_date;not null" json:"end_date"`                          // 结束日期
	Content   string    `gorm:"column:content;type:text;not null" json:"content"`                  // 回忆录内容文本
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`     // 创建日期
}

// TableName 指定表名
func (Memoir) TableName() string {
	return "memoirs"
}
