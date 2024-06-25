package http

import (
	"encoding/json"
	"errors"
	"log"
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

func (s *Server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *Server) respond(w http.ResponseWriter, data interface{}, status int, headers map[string]string) {
  var body interface{}
  switch v := data.(type) {
  case error:
    var errs []string
    for ce := v; ce != nil; ce = errors.Unwrap(ce) {
      errs = append(errs, ce.Error())
    }
    body = errs
  default:
    body = data
  }

  w.WriteHeader(status)
  if data != nil {
    err := json.NewEncoder(w).Encode(body)
    if err != nil {
      log.Fatalln(err)
    }
  }
}