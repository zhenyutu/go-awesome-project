package geecache

import (
	"awesomeProject/project/geecache/lru"
	"sync"
)

type Cache struct {
	mu         sync.Mutex
	cache      *lru.Cache
	cacheBytes int64
}

func (c *Cache) Put(key string, value CacheData) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cache == nil {
		c.cache = lru.New(c.cacheBytes)
	}
	c.cache.Put(key, &value)
}

func (c *Cache) Get(key string) (*CacheData, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cache == nil {
		return nil, false
	}

	value, ok := c.cache.Get(key)
	if ok {
		return value.(*CacheData), true
	}
	return nil, false
}
