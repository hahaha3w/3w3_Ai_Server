package domain

import (
	"context"
	"time"
)

type MemoirRepo interface {
	CreateMemoir(ctx context.Context, memoir *Memoir) (err error)
	GetUserMessages(ctx context.Context, userID int, startTime, endTime time.Time) (string, error)
	GetMemoirsByUserID(ctx context.Context, userID int, style, memoirType, startDate, endDate string, page, pageSize int32) (memoirs []*Memoir, total int32, err error)
	GetMemoirByID(ctx context.Context, memoirID, userID int) (memoir *Memoir, err error)
	DeleteMemoir(ctx context.Context, memoirID, userID int) (rows int64, err error)
}

type MemoirUsecase interface {
	GenerateMemoir(ctx context.Context, userID int, style, memoirType, startDate, endDate string) (*Memoir, error)
	GetMemoirList(ctx context.Context, userID int, memoirType, startDate, endDate string, page, pageSize int32) ([]*Memoir, bool, error)
	GetMemoirDetail(ctx context.Context, memoirID, userID int) (*Memoir, error)
	DeleteMemoir(ctx context.Context, memoirID, userID int) error
}
