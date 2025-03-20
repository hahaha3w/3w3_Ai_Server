package main

import (
	"github.com/cloudwego/kitex/server"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user/userservice"
	"github.com/hahaha3w/3w3_Ai_Server/user/internal/core"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"net"
)

func main() {
	godotenv.Load()
	err := core.LoadConfig()
	if err != nil {
		panic(err)
	}
	core.StartMtl()
	opts := kitexInit()
	userDelivery := wireApp()
	srv := userservice.NewServer(userDelivery, opts...)
	err = srv.Run()
	if err != nil {
		panic(err)
	}
}

func kitexInit() (opts []server.Option) {
	// address
	_, err := net.ResolveTCPAddr("tcp", viper.GetString("service.address"))
	if err != nil {
		panic(err)
	}

	//opts = append(opts,
	//	server.WithServiceAddr(addr),
	//	server.WithSuite(
	//		serversuite.CommonServerSuite{
	//			CurrentServiceName: viper.GetString("service.name"),
	//			RegistryAddr:       viper.GetString("consul.address")}))
	return
}
