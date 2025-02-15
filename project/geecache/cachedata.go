package geecache

type CacheData struct {
	data []byte
}

func (c *CacheData) Len() int64 {
	return int64(len(c.data))
}

func (c *CacheData) String() string {
	return string(c.data)
}
