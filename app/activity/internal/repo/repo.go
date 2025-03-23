package repo

import (
	"github.com/google/wire"
	"github.com/hahaha3w/3w3_Ai_Server/activity/internal/domain"
)

var ProviderSet = wire.NewSet(wire.Bind(new(domain.ActivityRepo), new(*MysqlActivityRepo)), NewMysqlActivityRepo)
