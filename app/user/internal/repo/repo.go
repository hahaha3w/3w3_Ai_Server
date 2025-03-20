package repo

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.UserRepo), new(*MysqlUserRepo)), NewMysqlUserRepo)
