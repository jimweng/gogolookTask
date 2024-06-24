package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	task "github.com/jimweng/gogolookTask"
)

type Server struct {
	router *chi.Mux
	svc    task.Service
}

func NewServer(taskSvc task.Service) *Server {
	s := Server{
		router: chi.NewRouter(),
		svc:    taskSvc,
	}
	s.routes()

	return &s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}