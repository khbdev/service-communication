package handler

import (
	"GeteWay/pkg/cache"
	"GeteWay/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToService(c *gin.Context, cCache *cache.Cache) {
	serviceName := c.Param("service")
	path := c.Param("path")

	// 🔹 1. Cache’dan status tekshirish
	status, ok := cCache.Get(serviceName)
	if !ok || !status.Health {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service not healthy"})
		return
	}

	// 🔹 2. Service URL olish
	baseURL, ok := service.Services[serviceName]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not registered"})
		return
	}

targetURL := baseURL + path + "?" + c.Request.URL.RawQuery

	// 🔹 3. Request yaratish
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}

	// 🔹 4. Headers nusxalash
	for k, v := range c.Request.Header {
		req.Header[k] = v
	}

	// 🔹 5. Requestni yuborish
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to reach service"})
		return
	}
	defer resp.Body.Close()

	// 🔹 6. Response qaytarish
	c.Status(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	c.Writer.Write(body)
}
