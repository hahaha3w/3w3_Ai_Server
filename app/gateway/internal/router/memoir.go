package router

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

func (r *Group) SetMemoirRouter() {
	group := r.Group("/memoir")
	var memoirApi api.MemoirApi

	groupAuthed := group.Use(middleware.JWT())
	groupAuthed.GET("/list", memoirApi.GetMemoirList)
	groupAuthed.GET("/:memoir_id", memoirApi.GetMemoirDetail)
	group.DELETE("/:memoir_id", memoirApi.DeleteMemoir)
}
