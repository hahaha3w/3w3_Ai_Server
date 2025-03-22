package domain

import (
	"time"
)

type Activity struct {
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建日期
}

// TableName 指定表名
func (Activity) TableName() string {
	return "activities"
}
