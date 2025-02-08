package main

import (
	"fmt"
	"time"
)

func returnsInt() int {
	defer func() {
		fmt.Println("Deferred function called")
	}()
	return 42
}

func fun1() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()

	// 规则二 defer执行顺序为先进后出

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()

	// 规则三 defer可以读取有名返回值（函数指定了返回参数名）
	return 100 //这里实际结果为2。如果是return 100呢
}

func fun2() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i) // 打印结果为 defer: 2
	}()

	defer func() {
		i++
		fmt.Println("defer1:", i) // 打印结果为 defer: 1
	}()
	return i
}

// 记录方法耗时
func functionCostTime(start time.Time) {
	cost := time.Since(start)
	fmt.Println("cost:", cost)
}

func main() {
	defer functionCostTime(time.Now())
	fmt.Println("Returned value:", returnsInt())

	fmt.Println(fun1())
	fmt.Println("**************************")

	fmt.Println(fun2())
}
