package service

import "log"

type User struct {
	Name string
	Age  int
}

type UserService struct {
}

func (u *UserService) Hello() {
	log.Println("hello world")
}

func (u *UserService) SayHello(name string) string {
	return "hello " + name
}

func (us *UserService) GetName(u *User) string {
	return u.Name
}

func (us *UserService) GetAge(u *User) int {
	return u.Age
}
