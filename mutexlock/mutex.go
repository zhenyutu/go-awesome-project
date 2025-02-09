package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

type atomicInt int

func (i *atomicInt) incr() {
	lock.Lock()
	defer lock.Unlock()

	*i++
}

func (i *atomicInt) get() {
	lock.Lock()
	defer lock.Unlock()

	fmt.Printf("%d\n", *i)
}

func main() {
	var a atomicInt
	a.incr()
	for i := 100; i > 0; i-- {
		go a.incr()
	}

	time.Sleep(1 * time.Second)
	a.get()
}
