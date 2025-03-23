package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat/chatservice"
	"log"
)

func main() {
	err := core.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	core.LoadLogger()
	delivery := wireApp()
	delivery.InitSubscribe()
	svr := chatservice.NewServer(delivery)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
