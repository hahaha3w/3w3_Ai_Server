package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat/echo"
	"log"
)

func main() {
	_ = core.LoadConfig()
	core.LoadLogger()
	delivery := wireApp()
	delivery.InitSubscribe()
	svr := echo.NewServer(delivery)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
