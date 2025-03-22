package delivery

import (
	"context"
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity"
)

var ProviderSet = wire.NewSet(New)

type ActivityDelivery struct {
	repo    domain.ActivityRepo
	usecase domain.ActivityUsecase
}

func New(repo domain.ActivityRepo, usecase domain.ActivityUsecase) *ActivityDelivery {
	return &ActivityDelivery{
		repo:    repo,
		usecase: usecase,
	}
}

func (d *ActivityDelivery) CreateUserActivity(ctx context.Context, req *activity.CreateUserActivityReq) (res *activity.CreateUserActivityResp, err error) {
	panic("implement me")
}

func (d *ActivityDelivery) GetUserActivities(ctx context.Context, req *activity.GetUserActivityReq) (res *activity.GetUserActivityResp, err error) {
	panic("implement me")
}
