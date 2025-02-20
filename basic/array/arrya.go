package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr := make([]int, 0)
	fmt.Println(arr)

	test(arr)
}

func test(i interface{}) {
	fmt.Printf(reflect.TypeOf(i).String(), "\n")
}
