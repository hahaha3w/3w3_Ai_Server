//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (


	"github.com/google/wire"
)

func wireApp() *delivery.UserDelivery {
	panic(wire.Build(
		core.ProviderSet,
		repo.ProviderSet,
		usecase.ProviderSet,
		delivery.ProviderSet,
	))
}
