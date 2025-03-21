package delivery

import "github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"

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
func (d Delivery) Get() {

}
