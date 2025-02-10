package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Dispatcher struct {
	handlers map[string]HandleFunc
}

func New() *Dispatcher {
	return &Dispatcher{handlers: make(map[string]HandleFunc)}
}

func (d *Dispatcher) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	d.handlers[key] = handler
}

func (d *Dispatcher) GET(pattern string, handler HandleFunc) {
	d.addRoute(http.MethodGet, pattern, handler)
}

func (d *Dispatcher) POST(pattern string, handler HandleFunc) {
	d.addRoute(http.MethodPost, pattern, handler)
}

// 实现Handler接口
func (dispatch *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := dispatch.handlers[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}

}

// 对外提供服务启动入口
func (dispatch *Dispatcher) Run(addr string) {
	http.ListenAndServe(addr, dispatch)
}
