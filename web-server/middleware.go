package webserver

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var flag bool = true
			fmt.Println("Checking authentication")
			if flag {
				hf(w, r)
			} else {
				return
			}
		}
	}
}
