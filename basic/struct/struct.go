package main

import "fmt"

type Human struct {
	name string
}

type Student struct {
	Human
	no uint
	string
}

func main() {
	var stu Student
	fmt.Println(stu)
	stu2 := Student{Human{"James"}, 111000, "good"}
	fmt.Println(stu2)
	fmt.Println(stu2.string)
}
