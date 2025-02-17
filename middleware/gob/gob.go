package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func EncodeTest() {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(Person{"Bob", 20})
	if err != nil {
		fmt.Println("Encode error:", err)
		return
	}
	fmt.Println("Encode result:", buffer.Bytes())
}

func DecodeTest() {
	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(Person{"Bob", 20})
	if err != nil {
		fmt.Println("Encode error:", err)
		return
	}

	decode := gob.NewDecoder(&buffer)
	var person Person
	err = decode.Decode(&person)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}
	fmt.Println("Decode result:", person)

}

func main() {
	EncodeTest()
	DecodeTest()
}
