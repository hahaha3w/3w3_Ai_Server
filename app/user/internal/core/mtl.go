package core

import (
	"github.com/spf13/viper"
)

func StartMtl() {
	mtl.InitTracing(viper.GetString("service.name"))
	mtl.InitMetric(viper.GetString("service.name"),
		viper.GetString("core.address"),
		viper.GetString("consul.address"))
}
