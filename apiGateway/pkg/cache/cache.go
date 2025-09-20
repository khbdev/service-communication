package cache

import "sync"

type ServiceStatus struct {
	Health bool
}

type Cache struct {
	data map[string]ServiceStatus
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{data: make(map[string]ServiceStatus)}
}

func (c *Cache) Set(key string, status ServiceStatus) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = status
}

func (c *Cache) Get(key string) (ServiceStatus, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	status, ok := c.data[key]
	return status, ok
}
