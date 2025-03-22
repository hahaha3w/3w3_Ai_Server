package delivery

import (
	"context"
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir"
)

var ProviderSet = wire.NewSet(New)

type MemoirDelivery struct {
	repo    domain.MemoirRepo
	usecase domain.Usecase
}

func New(repo domain.MemoirRepo, usecase domain.Usecase) *MemoirDelivery {
	return &MemoirDelivery{
		repo:    repo,
		usecase: usecase,
	}
}
func (d *MemoirDelivery) GenerateMemoir(ctx context.Context, req *memoir.GenerateDailyMemoirRequest) (res *memoir.GenerateMemoirResponse, err error) {
	//TODO : implement me
	return nil, err
}
func (d *MemoirDelivery) GetMemoirList(ctx context.Context, req *memoir.GetMemoirListRequest) (res *memoir.GetMemoirListResponse, err error) {
	//TODO : implement me
	return nil, err
}
func (d *MemoirDelivery) GetMemoirDetail(ctx context.Context, req *memoir.GetMemoirDetailRequest) (res *memoir.GetMemoirDetailResponse, err error) {
	//TODO : implement me
	return nil, err
}
