package router

import (
	"awesomeProject/project/gee/common"
	"fmt"
	"net/http"
)

/**
 * Router Definition
 */
type Router struct {
	handlers *Tire
}

func (r *Router) addRoute(method string, pattern string, handler common.HandleFunc) {
	key := method + "-" + pattern
	r.handlers.InsertKeyValue(key, handler)
}

/**
 * Dispatcher definition
 */
type Dispatcher struct {
	router *Router
}

func New() *Dispatcher {
	router := &Router{handlers: &Tire{}}
	return &Dispatcher{router: router}
}

func (d *Dispatcher) addRoute(method string, pattern string, handler common.HandleFunc) {
	d.router.addRoute(method, pattern, handler)
}

func (d *Dispatcher) GET(pattern string, handler common.HandleFunc) {
	d.addRoute(http.MethodGet, pattern, handler)
}

func (d *Dispatcher) POST(pattern string, handler common.HandleFunc) {
	d.addRoute(http.MethodPost, pattern, handler)
}

// 实现Handler接口
func (dispatch *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler := dispatch.router.handlers.Search(key); handler != nil {
		handler.Value(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// 对外提供服务启动入口
func (dispatch *Dispatcher) Run(addr string) {
	http.ListenAndServe(addr, dispatch)
}
