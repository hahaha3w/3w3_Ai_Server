package repo

import (
	"gorm.io/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}
func (r *ChatRepo) StoreChatRecord(content string, role string, containerID string) error {

	return nil
}
