package main

import (
	"fmt"
	"sync"
)

// ADD()
// WAIT()
// DONE()
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
