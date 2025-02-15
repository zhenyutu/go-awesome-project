package lru

import "container/list"

type Cache struct {
	maxSize int64
	used    int64
	list    *list.List
	cache   map[string]*list.Element
}

type Value interface {
	Len() int64
}

type entry struct {
	key   string
	value Value
}

func New(maxSize int64) *Cache {
	return &Cache{maxSize: maxSize, list: list.New(), cache: make(map[string]*list.Element)}
}

func (c *Cache) get(key string) (v Value, ok bool) {
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		kv := e.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}

func (c *Cache) put(key string, value Value) {
	if e, ok := c.cache[key]; ok {
		c.list.MoveToFront(e)
		kv := e.Value.(*entry)
		c.used = c.used + value.Len() - kv.value.Len()
		kv.value = value
	} else {
		kv := entry{key, value}
		c.cache[key] = c.list.PushFront(&kv)
		c.used = c.used + value.Len()
	}

	//删除最近未使用节点
	for c.maxSize > 0 && c.used > c.maxSize {
		ele := c.list.Back()
		if ele != nil {
			c.list.Remove(ele)
			kv := ele.Value.(*entry)
			delete(c.cache, kv.key)

			c.used = c.used - value.Len()
		}

	}
}
