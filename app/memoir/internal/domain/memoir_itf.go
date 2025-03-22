package domain

import (
	"context"
)

type MemoirRepo interface {
	CreateMemoir(ctx context.Context, memoir *Memoir) (err error)
	GetMemoirsByUserID(ctx context.Context, userID int, memoirType, startDate, endDate string, page, pageSize int32) (memoirs []*Memoir, total int32, err error)
	GetMemoirByID(ctx context.Context, memoirID, userID int) (memoir *Memoir, err error)
	DeleteMemoir(ctx context.Context, memoirID, userID int) (rows int64, err error)
}

type MemoirUsecase interface {
	GenerateMemoir(ctx context.Context, userID int, title, content, memoirType, startDate, endDate string) (*Memoir, error)
	GetMemoirList(ctx context.Context, userID int, memoirType, startDate, endDate string, page, pageSize int32) ([]*Memoir, int32, error)
	GetMemoirDetail(ctx context.Context, memoirID, userID int) (*Memoir, error)
	DeleteMemoir(ctx context.Context, memoirID, userID int) error
}
