package routerGroup

import (
	"fmt"
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
