package repo

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
	"gorm.io/gorm"
)

var _ domain.ActivityRepo = &MysqlActivityRepo{}

type MysqlActivityRepo struct {
	db *gorm.DB
}

func (m MysqlActivityRepo) CreateUserActivity(ctx context.Context, userId int64, relationId int64, ActivityType, description string) (activity *domain.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlActivityRepo) GetUserActivity(ctx context.Context, userId int64) (activity *[]domain.Activity, err error) {
	//TODO implement me
	panic("implement me")
}

func NewMysqlActivityRepo(db *gorm.DB) *MysqlActivityRepo {
	return &MysqlActivityRepo{
		db: db,
	}
}
