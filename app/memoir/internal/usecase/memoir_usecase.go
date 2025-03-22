package usecase

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"time"
)

var _ domain.MemoirUsecase = &ConcreteMemoirUsecase{}

type ConcreteMemoirUsecase struct {
	repo  domain.MemoirRepo
	cache *redis.Client
}

func NewConcreteMemoirUsecase(repo domain.MemoirRepo, cache *redis.Client) *ConcreteMemoirUsecase {
	return &ConcreteMemoirUsecase{
		repo:  repo,
		cache: cache,
	}
}

// GenerateMemoir 生成回忆录
func (c ConcreteMemoirUsecase) GenerateMemoir(ctx context.Context, userID int, title, content, memoirType, startDate, endDate string) (*domain.Memoir, error) {
	// 将字符串类型转化成 time.Time 类型
	const layout = "2006-01-02 15:04:05"

	var startTime, endTime time.Time
	var err error
	var now = time.Now()

	if startDate != "" && endDate != "" {
		startTime, err = time.Parse(layout, startDate)
		if err != nil {
			return nil, errors.New("invalid startDate format, expected YYYY-MM-DD")
		}

		endTime, err = time.Parse(layout, endDate)
		if err != nil {
			return nil, errors.New("invalid endDate format, expected YYYY-MM-DD")
		}

		now.Format(layout)
	}

	// 创建回忆录对象
	memoir := &domain.Memoir{
		UserID:    userID,
		Title:     title,
		Content:   content,
		Type:      memoirType,
		StartDate: startTime,
		EndDate:   endTime,
		CreatedAt: now,
	}

	// 调用 repo 层的 CreateMemoir 方法保存回忆录
	if err := c.repo.CreateMemoir(ctx, memoir); err != nil {
		return nil, err
	}

	// 返回创建的回忆录
	return memoir, nil
}

// GetMemoirList 获取用户的回忆录列表，支持分页、类型过滤和日期范围过滤
func (c ConcreteMemoirUsecase) GetMemoirList(ctx context.Context, userID int, memoirType, startDate, endDate string, page, pageSize int32) ([]*domain.Memoir, int32, error) {
	// 调用 repo 层的 GetMemoirsByUserID 方法获取回忆录列表
	memoirs, total, err := c.repo.GetMemoirsByUserID(ctx, userID, memoirType, startDate, endDate, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// 返回回忆录列表和总数
	return memoirs, total, nil
}

// GetMemoirDetail 获取某个回忆录的详细信息
func (c ConcreteMemoirUsecase) GetMemoirDetail(ctx context.Context, memoirID, userID int) (*domain.Memoir, error) {
	// 调用 repo 层的 GetMemoirByID 方法获取回忆录详情
	memoir, err := c.repo.GetMemoirByID(ctx, memoirID, userID)
	if err != nil {
		return nil, err
	}

	// 返回回忆录详情
	return memoir, nil
}

// DeleteMemoir 删除某个回忆录
func (c ConcreteMemoirUsecase) DeleteMemoir(ctx context.Context, memoirID, userID int) error {
	// 调用 repo 层的 DeleteMemoir 方法删除回忆录
	if err := c.repo.DeleteMemoir(ctx, memoirID, userID); err != nil {
		return err
	}

	return nil
}
