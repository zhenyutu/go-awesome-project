package client

import "testing"

func TestClientWithoutParams(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	err := client.invoke("UserService", "Hello", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestClientWithParams(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	err := client.invoke("UserService", "SayHello", []interface{}{"John"})
	if err != nil {
		t.Error(err)
	}
}
