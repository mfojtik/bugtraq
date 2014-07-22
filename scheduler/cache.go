package scheduler

// Credits: http://openmymind.net/Do-More-In-Process-Caching/

import (
	"sync"
)

type CacheItem interface {
	GetId() string
}

type Cache struct {
	items map[string]CacheItem
	lock  *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]CacheItem, 1024),
		lock:  new(sync.RWMutex),
	}
}

func (c *Cache) Get(id string) CacheItem {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.items[id]
}

func (c *Cache) Add(item CacheItem) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[item.GetId()] = item
}

func (c *Cache) Remove(id string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, id)
}
