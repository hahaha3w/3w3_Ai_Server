package repo

import (
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"gorm.io/gorm"
)

var _ domain.MemoirRepo = &MysqlMemoirRepo{}

type MysqlMemoirRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlMemoirRepo {
	return &MysqlMemoirRepo{
		db: db,
	}
}
