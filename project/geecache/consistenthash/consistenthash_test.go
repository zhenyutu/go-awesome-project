package consistenthash

import (
	"hash/crc32"
	"testing"
)

func TestConsistentHash(test *testing.T) {
	hr := New(3, crc32.ChecksumIEEE)
	hr.Add("8")
	hr.Add("18")
	hr.Add("28")

	hr.Get("7")

	//for i := 0; i < 100; i++ {
	//	test.Log(hr.Get(strconv.Itoa(i)))
	//}
}
