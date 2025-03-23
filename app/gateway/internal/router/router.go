package router

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	routerGroup := RouterGroup{router.Group("/api")}
	routerGroup.SetUserRouter()

	return r
}
