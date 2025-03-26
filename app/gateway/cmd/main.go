package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/router"
	"github.com/spf13/viper"
	"log"
)

func main() {
	err := core.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	r := router.InitRouter()
	rpc.InitClient()
	err = r.Run(viper.GetString("server.address"))
	if err != nil {
		log.Fatalf(err.Error())
	}
}
