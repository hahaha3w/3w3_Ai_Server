package domain

import (
	"time"
)

type Activity struct {
	ActivityID  int       `gorm:"column:activity_id;primaryKey;autoIncrement" json:"activity_id"`   // 活动唯一标识
	UserID      int       `gorm:"column:user_id;not null" json:"user_id"`                           // 关联用户ID
	RelatedID   int       `gorm:"column:related_id" json:"related_id"`                              // 关联资源ID
	Type        string    `gorm:"column:type;type:varchar(20);not null;default:''" json:"type"`     // 活动类型，chat, memoir, ai
	Description string    `gorm:"column:description;type:varchar(100);not null" json:"description"` // 活动描述
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`    // 发生时间
}

// TableName 指定表名
func (Activity) TableName() string {
	return "activities"
}
