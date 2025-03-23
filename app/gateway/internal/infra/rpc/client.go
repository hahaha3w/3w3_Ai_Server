package rpc

import (
	"log"

	"github.com/spf13/viper"

	"github.com/hahaha3w/3w3_Ai_Server/common/clientsuite"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user/userservice"

	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	UserClient   userservice.Client
	once         sync.Once
	err          error
	registryAddr string
	commonSuite  client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = viper.GetString("consul.address")
		serviceName := viper.GetString("service.name")
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		})
		initUserClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
