package geecache

import (
	"log"
	"testing"
)

func TestGroupCache(test *testing.T) {
	group := NewGroup("test", 1024, GetterFunc(func(key string) ([]byte, error) {
		log.Println("load db")
		return []byte("load by db"), nil
	}))
	group.Put("key1", CacheData{data: []byte("value1")})
	group.Put("key2", CacheData{data: []byte("value2")})
	group.Put("key3", CacheData{data: []byte("value3")})

	test.Log(group.Get("key1"))
	test.Log(group.Get("key4"))
	test.Log(group.Get("key4"))
}
