package routerGroup

import (
	"awesomeProject/project/gee/common"
	"awesomeProject/project/gee/router"
	"fmt"
	"net/http"
)

/**
 * Router Definition
 */
type Router struct {
	handlers *router.Tire
}

func (r *Router) addRoute(method string, pattern string, handler common.HandleFunc) {
	key := method + "-" + pattern
	r.handlers.InsertKeyValue(key, handler)
}

/**
 * Router Group Definition
 */
type RouterGroup struct {
	dispatcher  *Dispatcher
	prefix      string
	middlewares []common.HandleFunc
}

func (group *RouterGroup) addRoute(method string, comp string, handler common.HandleFunc) {
	pattern := group.prefix + comp
	group.dispatcher.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler common.HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler common.HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

/**
 * Dispatcher definition
 */
type Dispatcher struct {
	router *Router
	groups []*RouterGroup
}

func New() *Dispatcher {
	router := &Router{handlers: &router.Tire{}}
	groups := make([]*RouterGroup, 0)
	return &Dispatcher{router: router, groups: groups}
}

func (d *Dispatcher) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{prefix: prefix, dispatcher: d}
	d.groups = append(d.groups, newGroup)
	return newGroup
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
