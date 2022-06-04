package webserver

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola mundo")
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ruta api")
}
