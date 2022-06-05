package webserver

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	name  string
	email string
	phone string
}

type MetaData interface {
}
