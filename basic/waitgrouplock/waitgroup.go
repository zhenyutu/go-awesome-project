package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f() {
	fmt.Println("f call")
	wg.Done()
}

func main() {
	wg.Add(2)
	go f()
	go f()
	wg.Wait()
	fmt.Println("done")
}
