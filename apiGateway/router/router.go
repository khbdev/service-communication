// router/router.go
package router

import (
	"GeteWay/handler"
	"GeteWay/pkg/cache"
	cronhr "GeteWay/pkg/cronHR"

	"GeteWay/service"

	"github.com/gin-gonic/gin"
)

var CCache = cache.New() // global cache

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// start polling
	go cronhr.StartPolling(service.Services, CCache)

	r.Any("/api/:service/*path", func(c *gin.Context) {
		handler.ProxyToService(c, CCache)
	})	
	return r
}
