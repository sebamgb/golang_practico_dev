package webserver

import "net/http"

type Server struct {
	port   string
	router *Routes
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRoute(),
	}
}

func (s Server) Listen() (err error) {
	http.Handle("/", s.router)
	err = http.ListenAndServe(s.port, nil)
	if err != nil {
		return
	}
	return nil
}
