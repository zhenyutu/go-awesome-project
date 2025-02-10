package main

import (
	"fmt"
	"time"
)

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send end ", i)
	}
}

func receive(c <-chan int) {
	for {
		fmt.Println("receive readey ", <-c)
	}
}

func main() {
	c := make(chan int, 10)

	go send(c)
	go receive(c)
	time.Sleep(3 * time.Second)
	close(c)
}
