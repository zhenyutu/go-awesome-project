package main

import (
	"awesomeProject/middleware/protocal/student"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test := student.Student{
		Name:   "hello",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}

	data, err := proto.Marshal(&test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println("marshal result: ", string(data))

	newTest := &student.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println("unmarshal result: ", newTest)

	if test.Male != newTest.Male {
		log.Fatal("test male not equal")
	}
}
