package rpc

import (
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity/activityservice"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/chat/chatservice"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/memoir/memoirservice"
	"log"

	"github.com/spf13/viper"

	"github.com/hahaha3w/3w3_Ai_Server/common/clientsuite"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user/userservice"

	"sync"

	"github.com/cloudwego/kitex/client"
)

var (
	UserClient     userservice.Client
	MemoirClient   memoirservice.Client
	ActivityClient activityservice.Client
	ChatClient     chatservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
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
		initMemoirClient()
		initActivityClient()
		initChatClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func initMemoirClient() {
	MemoirClient, err = memoirservice.NewClient("memoir", commonSuite)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func initActivityClient() {
	ActivityClient, err = activityservice.NewClient("activity", commonSuite)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
func initChatClient() {
	ChatClient, err = chatservice.NewClient("chat", commonSuite)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
