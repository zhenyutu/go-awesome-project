package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Println("hello world")
}

func testReflectTypeValue() {
	var x float64 = 1.2345

	fmt.Println("==TypeOf==")
	t := reflect.TypeOf(x)
	fmt.Println("type: ", t)
	fmt.Println("kind:", t.Kind())

	fmt.Println("==ValueOf==")
	v := reflect.ValueOf(x)
	fmt.Println("value: ", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("string:", v.String())
	fmt.Println("value:", v.Interface())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	fmt.Println("===kind===")
	type MyInt int
	var m MyInt = 5
	v = reflect.ValueOf(m)
	fmt.Println("kind:", v.Kind()) // Kind() 返回底层的类型 int
	fmt.Println("type:", v.Type()) // Type() 返回类型 MyInt
}

func testObjectReflect() {
	user := &User{}
	t := reflect.TypeOf(user)
	fmt.Println("type:", t)
	v := reflect.ValueOf(user)
	fmt.Println("value type:", v.Type())
	fmt.Println("value type name:", v.Type().Name())
	fmt.Println("value type elem:", v.Type().Elem().String())

	iv := reflect.Indirect(v)
	fmt.Println("indirect value:", iv)
	fmt.Println("indirect value type:", iv.Type())
	fmt.Println("indirect value type name:", iv.Type().Name())
}

func testReflectMethod() {
	user := &User{"John", 18}
	t := reflect.ValueOf(user)
	fmt.Println("value:", t.IsValid())

	//Query
	filed := t.Elem().FieldByName("Age")
	fmt.Println("field: ", filed.IsValid(), filed)
	method := t.Elem().MethodByName("SayHello")
	fmt.Println("method: ", method.IsValid(), method)

	//Invoke
	method = t.MethodByName("SayHello")
	fmt.Println("method: ", method.IsValid(), method)
	result := method.Call([]reflect.Value{})
	fmt.Println("method call result: ", result)
}

func main() {
	//testReflectTypeValue()
	//testObjectReflect()
	testReflectMethod()
}
