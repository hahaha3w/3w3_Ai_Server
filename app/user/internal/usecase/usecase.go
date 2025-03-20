package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.UserUsecase), new(*ConcreteUserUsecase)), NewConcreteUserUsecase)
