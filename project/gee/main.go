package main

import (
	"awesomeProject/project/gee/gee"
	"fmt"
	"log"
	"net/http"
)

func handleFuncTest() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handler function call: ", r.URL.Path)
	})

	helloHander := func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	}
	http.HandleFunc("/hello", helloHander)

	http.ListenAndServe(":8080", nil)
}

type dispatcher struct{}

func (d *dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/hello":
		fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
	case "/test":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	case "/":
		fmt.Fprintf(w, r.URL.Path)
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func ListenHandleTest() {
	dispatcher := &dispatcher{}
	err := http.ListenAndServe(":8080", dispatcher)
	if err != nil {
		panic(err)
	}
}

func main() {

	handleFuncTest()
	ListenHandleTest()

	gee := gee.New()
	gee.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
	})
	gee.Run(":8080")
}
