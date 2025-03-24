package router

import (
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/api"
	"github.com/hahaha3w/3w3_Ai_Server/gateway/internal/middleware"
)

func (r *Group) SetMemoirRouter() {
	group := r.Group("/memoirs")
	var memoirApi api.MemoirApi

	groupAuthed := group.Use(middleware.JWT())
	groupAuthed.POST("/", memoirApi.GenerateMemoir)
	groupAuthed.GET("/list", memoirApi.GetMemoirList)
	groupAuthed.GET("/", memoirApi.GetMemoirDetail)
	group.DELETE("/", memoirApi.DeleteMemoir)
}
