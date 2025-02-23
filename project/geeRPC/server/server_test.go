package server

import (
	"awesomeProject/project/geeRPC/service"
	"reflect"
	"testing"
)

func TestServer_NewServer(t *testing.T) {
	//初始化服务
	server := NewServer("127.0.0.1:8080")
	//注册服务
	userService := service.UserService{}
	server.RegisterService(&userService)
	//启动服务
	err := server.Run()
	if err != nil {
		t.Errorf("server listen error: %v", err)
	}
	t.Log(server)
}

func TestServer_RegisterService(t *testing.T) {
	server := NewServer("127.0.0.1:8080")
	userService := &service.UserService{}
	server.RegisterService(userService)

	handler := server.services["UserService"].(*RPCHandler)
	object := handler.Object
	t.Log(reflect.TypeOf(object))
	method := object.MethodByName("SayHello")
	t.Log(method.IsValid())

	of := reflect.ValueOf(userService)
	method = of.MethodByName("SayHello")
	t.Log(method.IsValid())
}
