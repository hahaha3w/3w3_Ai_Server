package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/hahaha3w/3w3_Ai_Server/memoir/pkg/log"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/activity/activityservice"
	"sync"
)

var (
	activityClient activityservice.Client
	once           sync.Once
)

func ActivityClient() activityservice.Client {
	return activityClient
}

func InitRpc() {
	once.Do(func() {
		newClient, err := activityservice.NewClient("activityService", client.WithHostPorts("127.0.0.1:9011"))
		if err != nil {
			log.Log().Fatal(err)
		}
		activityClient = newClient
	})
}
