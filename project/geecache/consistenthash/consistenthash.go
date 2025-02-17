package consistenthash

import "sort"

type Hash func(data []byte) uint32

type HashRing struct {
	hash     Hash
	replicas int
	keys     []int
	hashmap  map[int]string
}

func New(replicas int, hash Hash) *HashRing {
	return &HashRing{
		hash:     hash,
		replicas: replicas,
		hashmap:  make(map[int]string),
	}
}

func (hm *HashRing) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < hm.replicas; i++ {
			hash := int(hm.hash([]byte(key)))
			hm.keys = append(hm.keys, hash)
			hm.hashmap[hash] = key
		}
	}

	sort.Ints(hm.keys)
}

func (hm *HashRing) Get(key string) string {
	hash := int(hm.hash([]byte(key)))

	idx := sort.Search(len(hm.keys), func(i int) bool {
		return hm.keys[i] >= hash
	})

	return hm.hashmap[hm.keys[idx%len(hm.keys)]]
}
