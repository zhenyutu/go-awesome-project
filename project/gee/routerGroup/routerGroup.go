package routerGroup

import (
	"awesomeProject/project/gee/common"
	"awesomeProject/project/gee/router"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func (group *RouterGroup) Use(middlewares ...common.HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

/**
 * Dispatcher definition
 */
type Dispatcher struct {
	*RouterGroup
	router *Router
	groups []*RouterGroup
}

func New() *Dispatcher {
	router := &Router{handlers: &router.Tire{}}
	dispatcher := &Dispatcher{router: router}

	dispatcher.RouterGroup = &RouterGroup{dispatcher: dispatcher}
	dispatcher.groups = []*RouterGroup{dispatcher.RouterGroup}

	return dispatcher
}

func Recovery() common.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Println("middleware default")
			if err := recover(); err != nil {
				http.Error(w, err.(string), http.StatusInternalServerError)
			}
		}()
		//此处recover未生效是因为defer和抛出异常的点必须在同一个函数内
		// throw panic error
	}
}

func Default() *Dispatcher {
	newDispatcher := New()
	newDispatcher.Use(Recovery())
	return newDispatcher
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
	var handlers []common.HandleFunc

	//此处defer recovery是生效的
	//defer func() {
	//	log.Println("hardcode default")
	//	if err := recover(); err != nil {
	//		http.Error(w, err.(string), http.StatusInternalServerError)
	//	}
	//}()

	for _, group := range dispatch.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			handlers = append(handlers, group.middlewares...)
		}
	}

	key := r.Method + "-" + r.URL.Path
	if handler := dispatch.router.handlers.Search(key); handler != nil {
		handlers = append(handlers, handler.Value)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
		return
	}

	for _, f := range handlers {
		f(w, r)
	}
}

// 对外提供服务启动入口
func (dispatch *Dispatcher) Run(addr string) {
	http.ListenAndServe(addr, dispatch)
}
