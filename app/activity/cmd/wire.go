//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/delivery"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/repo"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/usecase"
)

func wireApp() *delivery.ActivityDelivery {
	panic(wire.Build(
		core.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		delivery.ProviderSet,
	))
}
