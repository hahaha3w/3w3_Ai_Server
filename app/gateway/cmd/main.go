package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/router"
)

func main() {
	r := router.InitRouter()
	rpc.InitClient()
	err := r.Run(":8182")
	if err != nil {
		panic(err)
	}
}
