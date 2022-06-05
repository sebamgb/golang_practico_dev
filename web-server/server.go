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

func (s *Server) AddMiddleware(f http.HandlerFunc, middleawares ...Middleware) http.HandlerFunc {
	for _, m := range middleawares {
		f = m(f)
	}
	return f
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

func (s Server) Listen() (err error) {
	http.Handle("/", s.router)
	err = http.ListenAndServe(s.port, nil)
	if err != nil {
		return
	}
	return nil
}
