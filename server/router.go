package server

import "net/http"

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) {
	_, exists := r.rules[path]
	handler, methodExist := r.rules[path][method]

	return handler, exists, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
