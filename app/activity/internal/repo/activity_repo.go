package repo

import (
	"context"
	"errors"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
	"gorm.io/gorm"
)

var _ domain.ActivityRepo = &MysqlActivityRepo{}

type MysqlActivityRepo struct {
	db *gorm.DB
}

func NewMysqlActivityRepo(db *gorm.DB) *MysqlActivityRepo {
	return &MysqlActivityRepo{
		db: db,
	}
}

func (r MysqlActivityRepo) CreateUserActivity(ctx context.Context, activity *domain.Activity) (err error) {
	return r.db.WithContext(ctx).Model(&domain.Activity{}).Create(activity).Error
}

func (r MysqlActivityRepo) GetUserActivity(ctx context.Context, userId int64, page, pageSize int32) (activities []*domain.Activity, count int32, err error) {
	// 初始化查询
	query := r.db.WithContext(ctx).Model(&domain.Activity{}).Where("user_id = ?", userId)

	// 计算总数
	var totalCount int64
	if err = query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// 将 int64 转换为 int32，确保不会溢出
	if totalCount > int64(^uint32(0)>>1) { // 检查是否超出 int32 范围
		return nil, 0, errors.New("total count exceeds int32 limit")
	}
	count = int32(totalCount)

	// 分页查询
	offset := (page - 1) * pageSize
	if err = query.Offset(int(offset)).Limit(int(pageSize + 1)).Find(&activities).Order("created_at DESC").Error; err != nil {
		return nil, 0, err
	}

	return activities, count, nil
}
