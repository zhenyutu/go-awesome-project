package lru

import (
	"testing"
)

type String string

func (c String) Len() int64 { return int64(len(string(c))) }

func TestCache(test *testing.T) {
	cache := New(20)
	cache.put("key1", String("value1"))

	if v, ok := cache.get("key1"); !ok || v.(String) != "value1" {
		test.Fatal("cache put get check fail")
	}

}

func TestLru(t *testing.T) {
	cache := New(20)
	cache.put("key1", String("value1"))
	cache.put("key2", String("value2"))
	cache.put("key3", String("value3"))
	cache.put("key4", String("value4"))

	t.Log(cache.get("key1"))
	t.Log(cache.get("key2"))
	t.Log(cache.get("key3"))
	t.Log(cache.get("key4"))
}
