package router

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDispatcher(t *testing.T) {
	gee := New()

	gee.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
	})
	gee.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Path)
	})
	gee.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	gee.Run(":8080")
}
