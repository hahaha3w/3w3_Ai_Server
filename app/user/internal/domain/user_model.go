package domain

import (
	"time"
)

type User struct {
	UserID         int32  `gorm:"primary_key"`
	Email          string `gorm:"unique_index"`
	PasswordHashed string `gorm:"type:varchar(255) not null"`
	Nickname       string `gorm:"type:varchar(50)"`
	Avatar         string `gorm:"type:varchar(255)"`
	Phone          string `gorm:"type:varchar(20)"`
	Address        string `gorm:"type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (User) TableName() string {
	return "user"
}
