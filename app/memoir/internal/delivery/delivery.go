package delivery

import (
	"context"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
)

type Delivery struct {
	repo    domain.Repo
	usecase domain.Usecase
}

func New(repo domain.Repo, usecase domain.Usecase) *Delivery {
	return &Delivery{
		repo:    repo,
		usecase: usecase,
	}
}

func (d Delivery) GenerateDailyMemoir(ctx context.Context, req *memoir.GenerateDailyMemoirRequest) (res *memoir.GenerateMemoirResponse, err error) {
	//TODO : implement me
	return nil, err
}
func (d Delivery) GenerateWeeklyMemoir(ctx context.Context, req *memoir.GenerateWeeklyMemoirRequest) (res *memoir.GenerateMemoirResponse, err error) {
	//TODO : implement me
	return nil, err
}
func (d Delivery) GenerateMonthlyMemoir(ctx context.Context, req *memoir.GenerateMonthlyMemoirRequest) (res *memoir.GenerateMemoirResponse, err error) {
	//TODO : implement me
	return nil, err
}
func (d Delivery) GetMemoirList(ctx context.Context, req *memoir.GetMemoirListRequest) (res *memoir.GetMemoirListResponse, err error) {
	//TODO : implement me
	return nil, err
}
