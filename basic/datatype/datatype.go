package main

import "fmt"

func testByteArray() {
	a := []byte("world")
	b := []byte("hello")
	fmt.Println(a)
	fmt.Println(b)

	c := make([]byte, len(a)+len(b))
	copy(c[len(a):], b)
	copy(c, a)
	fmt.Println(c)
}

func main() {
	testByteArray()
}
