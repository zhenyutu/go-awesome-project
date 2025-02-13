package context

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	writer     http.ResponseWriter
	request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	context := Context{writer: w, request: r}
	context.Path = r.URL.Path
	context.Method = r.Method
	return &context
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.writer.WriteHeader(code)
}

func (c *Context) Header(key, value string) {
	c.writer.Header().Set(key, value)
}

func (c *Context) Data(code int, p []byte) {
	c.Status(code)
	_, err := c.writer.Write(p)
	if err != nil {
		http.Error(c.writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) String(code int, s string) {
	c.Header("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.writer.Write([]byte(s))
	if err != nil {
		http.Error(c.writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(code int, html string) {
	c.Header("Content-Type", "text/html")
	c.Status(code)
	_, err := c.writer.Write([]byte(html))
	if err != nil {
		http.Error(c.writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.writer, err.Error(), http.StatusInternalServerError)
	}
}

/*
* dispatcher definition
 */
type HandlerFunc func(*Context)

type Dispatcher struct {
	handlers map[string]HandlerFunc
}

func New() *Dispatcher {
	dispatcher := Dispatcher{make(map[string]HandlerFunc)}
	return &dispatcher
}

func (dispatcher *Dispatcher) AddHandler(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	dispatcher.handlers[key] = handler
}

func (dispatcher *Dispatcher) GET(pattern string, handler HandlerFunc) {
	dispatcher.AddHandler("GET", pattern, handler)
}

func (dispatcher *Dispatcher) POST(pattern string, handler HandlerFunc) {
	dispatcher.AddHandler("POST", pattern, handler)
}

func (dispatcher *Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	handler := dispatcher.handlers[key]
	if handler == nil {
		http.NotFound(w, r)
		return
	}
	handler(NewContext(w, r))
}

func (dispatcher *Dispatcher) Run() {
	http.ListenAndServe(":8080", dispatcher)
}
