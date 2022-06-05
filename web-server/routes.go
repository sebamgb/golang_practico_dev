package webserver

import (
	"net/http"
)

type Routes struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRoute() *Routes {
	return &Routes{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Routes) FindHandlers(path, method string) (handler http.HandlerFunc, methodExist bool, exist bool) {
	_, exist = r.rules[path]
	handler, methodExist = r.rules[path][method]
	return
}

func (r *Routes) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handle, methodexist, exists := r.FindHandlers(request.URL.Path, request.Method)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodexist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handle(w, request)
}
