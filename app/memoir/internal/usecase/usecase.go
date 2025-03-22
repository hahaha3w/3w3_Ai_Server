package usecase

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/internal/domain"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.MemoirUsecase), new(*ConcreteMemoirUsecase)), NewConcreteMemoirUsecase)
