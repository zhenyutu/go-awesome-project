package geecache

import (
	"sync"
)

/**
 * Getter
 */
type Getter interface {
	Get(key string) ([]byte, error)
}
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

/**
 * Group
 */
type Group struct {
	name   string
	getter Getter
	cache  Cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, maxSize int64, getter Getter) *Group {
	if getter == nil {
		panic("geecache: getter is nil")
	}

	mu.RLock()
	defer mu.RUnlock()

	group := Group{
		name:   name,
		getter: getter,
		cache:  Cache{cacheBytes: maxSize},
	}
	groups[name] = &group
	return &group
}

func getGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()

	return groups[name]
}

func (g *Group) Get(key string) (*CacheData, bool) {
	if v, ok := g.cache.Get(key); ok {
		return v, true
	}

	data, err := g.load(key)
	if err != nil {
		return nil, false
	}
	return data, true
}

func (g *Group) load(key string) (*CacheData, error) {
	data, err := g.getter.Get(key)
	if err != nil {
		return nil, err
	}

	cacheData := &CacheData{data: data}
	g.cache.Put(key, *cacheData)
	return cacheData, nil
}

func (g *Group) Put(key string, value CacheData) {
	g.cache.Put(key, value)
}
