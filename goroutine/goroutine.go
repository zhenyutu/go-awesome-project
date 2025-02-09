package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("hello world, %d", i)
			fmt.Println()
		}(i)
	}
}

func f2() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				arr[ii]++
				fmt.Printf("i:%d", ii)
				fmt.Println()
				runtime.Gosched()
			}

		}(i)
	}
	fmt.Println(arr)
}

func main() {
	defer calCost(time.Now())

	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				arr[ii]++
				//fmt.Printf("i:%d", ii)
				//fmt.Println()
				//runtime.Gosched()
			}

		}(i)
	}
	fmt.Println(runtime.NumGoroutine()) // 当前正在运行的goroutine 数
	fmt.Println(arr)
	time.Sleep(1 * time.Second)
	fmt.Println(arr)
	fmt.Println(runtime.NumCPU())

	f1()
	f2()
}

func calCost(start time.Time) {
	fmt.Println(time.Since(start))
}
