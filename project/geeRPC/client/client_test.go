package client

import "testing"

func TestClient(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	err := client.invoke("UserService", "Hello", "j")
	if err != nil {
		t.Error(err)
	}
}

func TestClient2(t *testing.T) {
	client := NewClient("127.0.0.1:8080")
	err := client.invoke("UserService", "SayHello", "john")
	if err != nil {
		t.Error(err)
	}
}
