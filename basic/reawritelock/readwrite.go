package main

import (
	"fmt"
	"sync"
	"time"
)

func read(i int, wg *sync.WaitGroup, m *sync.RWMutex) {
	m.RLock()
	defer m.RUnlock()

	fmt.Printf("read:%d", i)
	time.Sleep(1 * time.Second)

	wg.Done()
}

func write(i int, wg *sync.WaitGroup, m *sync.RWMutex) {
	m.Lock()
	defer m.Unlock()

	i++
	fmt.Println("write")
	time.Sleep(2 * time.Second)

	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}

	var data int
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go read(data, &wg, &mu)
		go write(data, &wg, &mu)
	}

	wg.Wait()
}
