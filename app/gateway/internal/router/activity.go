package router

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

func (r *Group) SetActivityRouter() {
	group := r.Group("/activities")
	var activityApi api.ActivityApi

	// 需要认证的路由
	groupAuthed := group.Use(middleware.JWT())
	groupAuthed.GET("/", activityApi.GetUserActivities)
}
