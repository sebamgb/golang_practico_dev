package webserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	query, exists := r.URL.Query()["nombre"]
	if !exists || len(query) <= 0{
		fmt.Fprintf(w, "no proporcionas nombre??")
	} else {
		fmt.Fprintf(w, "hola, "+query[0])
	}
}

func HandleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ruta api")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "aplication/json")
	log.Println(user.Name)
	w.Write(response)
}
