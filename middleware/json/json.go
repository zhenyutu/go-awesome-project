package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func EncodeTest() {
	person := Person{"John", 20}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Encode Fail:", err)
	}
	fmt.Println("Encode result: ", jsonBytes)
}

func DecodeTest() {
	person := Person{"John", 20}
	jsonBytes, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Encode Fail:", err)
	}

	var decodePerson Person
	err2 := json.Unmarshal(jsonBytes, &decodePerson)
	if err != nil {
		fmt.Println("Decode Fail:", err2)
	}
	fmt.Println("Decode result: ", decodePerson)
}

func main() {
	EncodeTest()
	DecodeTest()
}
