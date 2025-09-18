package router

import (
	"GeteWay/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	
	r.GET("/api/:service/*path", handler.ProxyToService)
	
	return r
}
