package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	task "github.com/jimweng/gogolookTask"
)

func (s *Server) HandleGetTasks() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
   res, err := s.svc.GetTasks()
   if err != nil {
    s.respond(w, err ,http.StatusBadRequest, nil)
    return
   }
   s.respond(w, res, http.StatusOK, nil)
  }
}

func (s *Server) HandlePosTask() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    var req task.Task
    err := s.decode(r, &req)
    if err != nil {
      s.respond(w, err, http.StatusBadRequest, nil)
    }

    res, err := s.svc.CreateTask(&task.Task{
      ID: uuid.New().String(),
      Name: req.Name,
      Status: req.Status,
    })

    if err != nil {
      s.respond(w, err, http.StatusInternalServerError, nil)
      return
    }

    s.respond(w, res, http.StatusCreated, nil)
  }
}

func (s *Server) HandleUpdateTask() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    var req task.Task
    err := s.decode(r, &req)
    if err != nil {
      s.respond(w, err, http.StatusBadRequest, nil)
      return
    }

    err = s.svc.UpdateTask(&task.Task{
      ID: id,
      Name: req.Name,
      Status: req.Status,
    })

    if err != nil {
      s.respond(w, err, http.StatusInternalServerError, nil)
      return
    }

    s.respond(w, nil, http.StatusOK, nil)
  }
}

func (s *Server) HandleDeleteTask() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    err := s.svc.DeleteTask(id)
    if err != nil {
      s.respond(w, err, http.StatusInternalServerError, nil)
      return
    }

    s.respond(w, nil, http.StatusOK, nil)
  }
}

func (s *Server) HandleGetTaskByID() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
   id := chi.URLParam(r, "id")
   res, err := s.svc.GetTaskByID(id)
   if err != nil {
    s.respond(w, err ,http.StatusBadRequest, nil)
    return
   }
   s.respond(w, res, http.StatusOK, nil)
  }
}