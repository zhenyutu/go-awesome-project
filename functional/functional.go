package main

import (
	"fmt"
	"reflect"
	"time"
)

type funcType func(time time.Time)

func GetTime(time time.Time) {
	fmt.Println(time)
}

func main() {
	x, y := swap(1, 2)
	fmt.Println(x, y)

	f := func() int { return 7 }
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(f())

	//直接使用GetTime
	GetTime(time.Now())

	//定义函数类型funcType变量f1
	var f1 funcType = GetTime
	f1(time.Now())

	//先把GetTime函数转为funcType类型，然后传入参数调用
	funcType(GetTime)(time.Now())

	//匿名函数
	sum := 0
	func() {
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Println(sum)
	}()
	fmt.Println(sum)
	fmt.Println("*************")

	//匿名函数值传递
	var sum2 = 0
	func(sum2 int) {
		for i := 1; i <= 100; i++ {
			sum2 += i
		}
		fmt.Println(sum2)
	}(sum2)
	fmt.Println(sum2)
	fmt.Println("*************")

	ff := closure()
	ff()
	ff()

}

func closure() func() int {
	var sum3 = 0
	f := func() int {
		for i := 1; i <= 100; i++ {
			sum3 += i
		}
		fmt.Println(sum3)
		return sum3
	}

	//只会在初始化时输出一次
	fmt.Println(sum3)
	return f
}

func swap(x int, y int) (t1 int, t2 int) {
	x, y = y, x
	return x, y
}

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}
