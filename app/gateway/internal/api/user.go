package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/infra/rpc"
	"github.com/hahaha3w/3w3_Ai_Server/rpc-gen/user"
)

type UserApi struct {
}

func (api *UserApi) Login(c *gin.Context) {

	rpc.UserClient.GetUserInfo(c, &user.GetUserInfoReq{})

}
func (api *UserApi) Register(c *gin.Context) {
}
