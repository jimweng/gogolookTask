package http

func (s *Server) routes() {
  s.router.Get("/tasks", s.HandleGetTasks())
  s.router.Post("/task", s.HandlePosTask())
  s.router.Put("/tasks/{id}", s.HandleUpdateTask())
  s.router.Delete("/tasks/{id}", s.HandleDeleteTask())
  s.router.Get("/task/{id}", s.HandleGetTaskByID())
}