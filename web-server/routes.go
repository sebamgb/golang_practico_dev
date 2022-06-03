package webserver

import (
	"fmt"
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

func (r *Routes) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "hello world")
}
