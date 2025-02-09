package main

import (
	"fmt"
	"time"
)

func f1(ch chan<- int) {
	for {
		time.Sleep(2 * time.Second)
		ch <- 1
	}

}

func f2(ch chan<- int) {
	for {
		time.Sleep(3 * time.Second)
		ch <- 2
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go f1(ch1)
	go f2(ch2)

	for {
		select {
		case i := <-ch1:
			fmt.Println("ch1 is ", i)
		case i := <-ch2:
			fmt.Println("ch2 is ", i)
		}
	}

}
