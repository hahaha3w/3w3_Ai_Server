package main

import (
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat/echo"
	"log"
)

func main() {
	_ = core.LoadConfig()
	svr := echo.NewServer(wireApp())
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
