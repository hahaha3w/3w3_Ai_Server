package main

import (
	"net"

	"github.com/joho/godotenv"
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
	addr, err := net.ResolveTCPAddr("tcp", viper.GetString("service.address"))
	if err != nil {
		panic(err)
	}

	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithSuite(
			serversuite.CommonServerSuite{
				CurrentServiceName: viper.GetString("service.name"),
				RegistryAddr:       viper.GetString("consul.address")}))
	return
}
