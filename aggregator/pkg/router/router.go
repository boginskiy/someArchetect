package router

import (
	"net/http"
)

type Router interface {
	Handle(method, path string, handler func(http.ResponseWriter, *http.Request))
	http.Handler
}

type Route struct {
	routes map[string]http.HandlerFunc
}

func NewRoute() *Route {
	return &Route{}
}

func (rt *Route) Handle(method, path string, handler func(http.ResponseWriter, *http.Request)) {
	if rt.routes == nil {
		rt.routes = make(map[string]http.HandlerFunc, 10)
	}
	rt.routes[method+path] = handler
}

func (rt *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, exists := rt.routes[r.Method+r.URL.Path]
	if !exists {
		http.NotFound(w, r)
		return
	}
	handler(w, r)
}
