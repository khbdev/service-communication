package handler

import (
	"GeteWay/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProxyToService(c *gin.Context) {
	serviceName := c.Param("service")
	path := c.Param("path") 


	baseURL, ok := service.Services[serviceName]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	
	targetURL := baseURL + path

	
	req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create request"})
		return
	}


	for k, v := range c.Request.Header {
		req.Header[k] = v
	}


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "failed to reach service"})
		return
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)
	c.Writer.Write(body)
}
