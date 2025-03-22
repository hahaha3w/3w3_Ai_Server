package domain

import "context"

type ActivityRepo interface {
	CreateUserActivity(ctx context.Context, userId int64, relationId int64, ActivityType, description string) (activity *Activity, err error)
	GetUserActivity(ctx context.Context, userId int64) (activity *[]Activity, err error)
}

type ActivityUsecase interface {
	CreateUserActivity(ctx context.Context, userId int64, relationId int64, ActivityType, description string) (activity *Activity, err error)
	GetUserActivity(ctx context.Context, userId int64) (activity *[]Activity, err error)
}
