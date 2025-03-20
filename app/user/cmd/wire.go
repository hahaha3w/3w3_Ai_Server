//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/delivery"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/repo"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/usecase"
)

func wireApp() *delivery.UserDelivery {
	panic(wire.Build(
		core.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		delivery.ProviderSet,
	))
}
