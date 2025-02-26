package client

import (
	"awesomeProject/project/geeRPC/example"
	"fmt"
	"testing"
)

func TestClientWithoutParams(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	data, err := client.invoke("UserService", "Hello", nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

func TestClientWithParams(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	data, err := client.invoke("UserService", "SayHello", []interface{}{"John"})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data))
}

func TestClientStub(test *testing.T) {
	client := NewClient("127.0.0.1:8080")

	userServiceClient := &example.UserServiceClient{}
	client.InitServiceStub(userServiceClient)

	res := userServiceClient.SayHello("John")
	fmt.Println(res)
}
