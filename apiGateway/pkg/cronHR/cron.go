package cronhr

import (
	"GeteWay/pkg/cache"
	"fmt"
	"net/http"
	"time"
)

func StartPolling(services map[string]string, c *cache.Cache) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for name, url := range services {
			status := checkService(url)
			c.Set(name, status)
			fmt.Printf("Updated cache for %s: %+v\n", name, status)
		}
	}
}

// 🔹 Ready endpoint olib tashlandi, faqat /health tekshiradi
func checkService(baseURL string) cache.ServiceStatus {
	status := cache.ServiceStatus{Health: false}

	resp, err := http.Get(baseURL + "/health")
	if err == nil && resp.StatusCode == 200 {
		status.Health = true
	}

	return status
}
