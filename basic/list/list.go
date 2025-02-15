package main

import (
	"container/list"
	"fmt"
)

func main() {

	list := list.New()
	v4 := list.PushBack(4)
	v2 := list.PushFront(2)

	list.InsertBefore(1, v2)
	list.InsertBefore(3, v4)

	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}
