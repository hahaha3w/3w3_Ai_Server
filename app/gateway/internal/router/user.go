package router

import "github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"

func (r *RouterGroup) SetUserRouter() {
	group := r.Group("/user")
	var userApi api.UserApi
	group.POST("/login", userApi.Login)
	group.POST("/register", userApi.Register)
}
