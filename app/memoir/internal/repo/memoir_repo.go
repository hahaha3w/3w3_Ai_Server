package repo

import (
	"context"
	"errors"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"gorm.io/gorm"
)

var _ domain.MemoirRepo = &MysqlMemoirRepo{}

type MysqlMemoirRepo struct {
	db *gorm.DB
}

func NewMysqlMemoirRepo(db *gorm.DB) *MysqlMemoirRepo {
	return &MysqlMemoirRepo{
		db: db,
	}
}

func (r MysqlMemoirRepo) CreateMemoir(ctx context.Context, memoir *domain.Memoir) (err error) {
	return r.db.WithContext(ctx).Model(&domain.Memoir{}).Create(memoir).Error
}

func (r MysqlMemoirRepo) GetMemoirsByUserID(ctx context.Context, userID int, memoirType, style, startDate, endDate string, page, pageSize int32) (memoirs []*domain.Memoir, total int32, err error) {
	// 初始化查询
	query := r.db.WithContext(ctx).Model(&domain.Memoir{}).Where("user_id = ?", userID)

	// 添加回忆录类型过滤条件
	if memoirType != "" {
		query = query.Where("type = ?", memoirType)
	}

	// 添加回忆录风格过滤条件
	if style != "" {
		query = query.Where("style = ?", style)
	}

	// 添加日期范围过滤条件
	if startDate != "" && endDate != "" {
		query = query.Where("start_date >= ? AND end_date <= ?", startDate, endDate)
	}

	// 计算总数
	var totalCount int64
	if err = query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// 将 int64 转换为 int32，确保不会溢出
	if totalCount > int64(^uint32(0)>>1) { // 检查是否超出 int32 范围
		return nil, 0, errors.New("total count exceeds int32 limit")
	}
	total = int32(totalCount)

	// 分页查询
	offset := (page - 1) * pageSize
	if err = query.Offset(int(offset)).Limit(int(pageSize)).Find(&memoirs).Error; err != nil {
		return nil, 0, err
	}

	return memoirs, total, nil
}

func (r MysqlMemoirRepo) GetMemoirByID(ctx context.Context, memoirID, userID int) (memoir *domain.Memoir, err error) {
	memoir = &domain.Memoir{}
	err = r.db.WithContext(ctx).Model(&domain.Memoir{}).
		Where("memoir_id = ? AND user_id = ?", memoirID, userID).
		First(memoir).Error
	return memoir, err
}

func (r MysqlMemoirRepo) DeleteMemoir(ctx context.Context, memoirID, userID int) (rows int64, err error) {
	result := r.db.WithContext(ctx).Model(&domain.Memoir{}).
		Where("memoir_id = ? AND user_id = ?", memoirID, userID).
		Delete(&domain.Memoir{})

	return result.RowsAffected, result.Error
}
