package server

import "net/http"

type Server struct {
	url    string
	router *Router
}

func NewServer(url string) *Server {
	return &Server{
		url:    url,
		router: NewRouter(),
	}
}

func (s *Server) RunServer() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.url, nil)

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exists := s.router.rules[path]

	if !exists {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}

	return f
}
