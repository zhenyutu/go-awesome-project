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
	fmt.Println("result: ", result)

	var result2 Result
	asyncCall := client.Go("Cal.CalSquare", 3, &result2, nil)
	<-asyncCall.Done
	fmt.Println("result2: ", result2)
}
