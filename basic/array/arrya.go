package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {

	arr := make([]byte, 30)
	fmt.Println(arr)

	cur := arr
	cur = cur[20:]
	bytes, _ := json.Marshal("hello john")
	fmt.Println(len(bytes))
	copy(cur, bytes)
	fmt.Println(arr)
}

func test(i interface{}) {
	fmt.Printf(reflect.TypeOf(i).String(), "\n")
}
