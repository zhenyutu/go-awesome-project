package main

import (
	"fmt"
	"reflect"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println(reflect.TypeOf(r))
			if err, ok := r.(error); !ok {
				err = fmt.Errorf("pkg: %v", r)
				fmt.Println(err)
				fmt.Println(reflect.TypeOf(err))
			}
		}
	}()

	fmt.Println("Recovered start")
	//var a, b = 1, 0
	//fmt.Println(a / b)
	panic("panic error")
	fmt.Println("Recovered end")
}

func main() {
	f()
}
