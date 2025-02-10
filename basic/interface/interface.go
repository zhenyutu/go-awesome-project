package main

import "fmt"

type I interface {
	f()
}

type T string

func (t T) f() {
	fmt.Println(string(t))
}

type S int

func (s S) f() {
	fmt.Println(s)
}

func main() {
	t := T("hello")
	t.f()

	var i I = T("world")
	i.f()

	var j I = S(10)
	j.f()

	//断言，判断接口I的实现是否为T类型
	if v, ok := i.(T); ok {
		fmt.Println("varI类型断言结果为：", v) // varI已经转为T类型
	}

	switch v := j.(type) {
	case T:
		fmt.Println("varJ类型断言结果为T", v) // varI已经转为T类型
	case S:
		fmt.Println("varJ类型断言结果为S", v) // varI已经转为S类型
	}
}
