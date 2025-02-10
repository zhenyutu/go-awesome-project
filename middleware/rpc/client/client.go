package main

import (
	"fmt"
	"net/rpc"
)

type Result struct {
	Num, Res int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	var result Result
	if err := client.Call("Cal.CalSquare", 2, &result); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
