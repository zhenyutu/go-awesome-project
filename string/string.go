package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "hello world\nbob"
	fmt.Println(s)

	s1 := "yes我爱中国"
	fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))

	s2 := strings.Join([]string{"hello", "world"}, ", ")
	fmt.Println(s2)

	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString("\t")
	buffer.WriteString(s2)
	fmt.Println(buffer.String())

	var buffer2 strings.Builder
	buffer2.WriteString(s1)
	buffer2.WriteString("\t")
	buffer2.WriteString(s2)
	fmt.Println(buffer2.String())

	var s3 = new(string)
	s3 = &s1
	fmt.Println(s3)
	fmt.Println(*s3)
	fmt.Println(reflect.TypeOf(s3))

	s4 := [5]string{"1"}
	s5 := [5]string{"1"}
	fmt.Println(s4 == s5)

	sayHello("john", "bob")

}

func sayHello(who ...string) {
	for _, w := range who {
		fmt.Printf("hello %s\n", w)
		fmt.Println()
	}
}
