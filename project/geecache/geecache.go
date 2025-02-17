package geecache

import (
	"log"
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
	picker PeerPicker

	//single run
	mu sync.RWMutex
	rm map[string]*run
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

/**
 * Single Run
 */
type run struct {
	wg    *sync.WaitGroup
	value interface{}
	err   error
}

func (g *Group) Run(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.rm == nil {
		g.rm = make(map[string]*run)
	}

	if runner, ok := g.rm[key]; ok {
		runner.wg.Wait()
		return runner.value, runner.err
	}

	runner := &run{wg: &sync.WaitGroup{}}
	runner.wg.Add(1)
	g.rm[key] = runner

	runner.value, runner.err = fn()
	runner.wg.Done()

	return runner.value, runner.err
}

/**
 * Data Query
 */
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
	data, err := g.Run(key, func() (interface{}, error) {
		if g.picker != nil {
			peer, _ := g.picker.Pick(key)
			if peer != nil {
				data, err := peer.PeerGet(g.name, key)
				if data != nil {
					cacheData := &CacheData{data: data}
					return cacheData, nil
				}
				log.Println("[GeeCache] Failed to get from peer", err)
			}
		}

		return g.loadLocally(key)
	})
	if err != nil {
		return nil, err
	}

	return data.(*CacheData), err
}

func (g *Group) loadLocally(key string) (*CacheData, error) {
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
