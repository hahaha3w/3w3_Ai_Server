package delivery

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserDelivery)
