package webserver

import (
	"net/http"
)

type Routes struct {
	rules map[string]http.HandlerFunc
}

func NewRoute() *Routes {
	return &Routes{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Routes) FindHandlers(path string) (handler http.HandlerFunc, exist bool) {
	handler, exist = r.rules[path]
	return
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handle, exists := r.FindHandlers(request.URL.Path)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handle(w, request)
}
