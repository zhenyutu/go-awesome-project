package routerGroup

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestRouterGroup(t *testing.T) {
	gee := New()
	group1 := gee.Group("/v1")
	group1.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "v1, %s!", r.FormValue("name"))
	})

	group2 := gee.Group("/v2")
	group2.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "v2, %s!", r.FormValue("name"))
	})

	gee.Run(":8080")
}

func TestRouterGroupMiddleware(t *testing.T) {
	gee := New()
	group1 := gee.Group("/v1")
	group1.Use(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 1")
	})
	group1.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "v1, %s!", r.FormValue("name"))
	})

	gee.Run(":8080")
}

func TestRouterGroupMiddlewareDefault(t *testing.T) {
	gee := Default()
	//gee := New()
	group1 := gee.Group("/v1")
	group1.Use(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware 1")
	})
	group1.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "v1, %s!", r.FormValue("name"))
	})
	group1.GET("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	gee.Run(":8080")
}
