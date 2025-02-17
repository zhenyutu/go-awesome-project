package main

import (
	"fmt"
	"sort"
)

// 查找有序数字中第一个满足条件的数
func search() {
	arr := []int{1, 3, 5, 7, 9, 10}
	res := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 6
	})
	fmt.Println(res)
}

func main() {
	search()
}
