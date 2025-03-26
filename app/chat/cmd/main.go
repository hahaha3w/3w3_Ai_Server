package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/hahaha3w/3w3_Ai_Server/chat/internal/core"
	"github.com/hahaha3w/3w3_Ai_Server/common/serversuite"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat/chatservice"
	"github.com/spf13/viper"
	"net"
)

func main() {
	core.LoadLogger()
	if err := core.LoadConfig(); err != nil {
		panic(err)
	}

	opts := kitexInit()
	chatDelivery := wireApp()
	svr := chatservice.NewServer(chatDelivery, opts...)
	if err := svr.Run(); err != nil {
		panic(err)
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", viper.GetString("service.address"))
	if err != nil {
		panic(err)
	}

	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts, server.WithSuite(
		serversuite.CommonServerSuite{
			CurrentServiceName: viper.GetString("service.name"),
			RegistryAddr:       viper.GetString("consul.address")}),
	)
	return
}
