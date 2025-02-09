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

func f3() {
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
}

func f4() {
	for {
		fmt.Println("call f4 function")
	}
}

func f5() {
	fmt.Println("call f5 function")
}

func main() {

	f4()
	f5()
	ch := make(chan int)
	<-ch

	f1()
	f2()
	f3()
}

func calCost(start time.Time) {
	fmt.Println(time.Since(start))
}
