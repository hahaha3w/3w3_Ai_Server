package domain

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
)

type ActivityRepo interface {
	CreateUserActivity(ctx context.Context, activity *Activity) (err error)
	GetUserActivity(ctx context.Context, userId int64, page, pageSize int32) (activity *[]Activity, count int32, err error)
}

type ActivityUsecase interface {
	CreateUserActivity(ctx context.Context, userId int64, relationId int64, ActivityType, description string) (resp *activity.CreateUserActivityResp, err error)
	GetUserActivity(ctx context.Context, userId int64, page, pageSize int32) (resp *activity.GetUserActivityResp, err error)
}
