package domain

import (
	"time"
)

type User struct {
	UserID         int32     `gorm:"primary_key;auto_increment" json:"user_id"`            // 用户唯一标识
	Username       string    `gorm:"type:varchar(50);not null" json:"username"`            // 用户名
	Email          string    `gorm:"type:varchar(100);unique_index;not null" json:"email"` // 邮箱
	PasswordHashed string    `gorm:"type:varchar(255);not null" json:"password"`           // 密码哈希值
	Avatar         string    `gorm:"type:varchar(255)" json:"avatar_url"`                  // 头像URL
	Bio            string    `gorm:"type:varchar(255)" json:"bio"`                         // 用户简介
	RegisterDate   time.Time `gorm:"type:datetime;not null" json:"register_date"`          // 注册时间
	Theme          string    `gorm:"type:varchar(20);default:'light'" json:"theme"`        // 主题偏好
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`                     // 创建时间
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`                     // 更新时间
}

func (User) TableName() string {
	return "user"
}
