package example

import "log"

type UserService struct {
}

func (u *UserService) Hello() {
	log.Println("hello world")
}

func (u *UserService) SayHello(name string) string {
	return "hello " + name
}
