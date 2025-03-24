package router

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

func (r *Group) SetUserRouter() {
	group := r.Group("/user")
	var userApi api.UserApi
	group.POST("/send-code", userApi.SendCode)
	group.POST("/login", userApi.Login)
	group.POST("/register", userApi.Register)

	groupAuthed := group.Use(middleware.JWT())
	groupAuthed.PUT("/info", userApi.UpdateUserInfo)
	groupAuthed.PUT("/password", userApi.ChangePassword)
	groupAuthed.GET("/info", userApi.GetUserInfo)
	group.DELETE("/", userApi.DeleteUser)
}
