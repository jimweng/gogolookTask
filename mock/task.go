package mock

import gtask "github.com/jimweng/gogolookTask"

// This ensures the Service type implements the Service interface via a compiler check,
// even if it is not used elsewhere.  You can read more on this pattern on the effective go site
// https://golang.org/doc/effective_go#blank_implements
var _ gtask.Service = (*Service)(nil)

type Service struct {
	CreateTaskFn  func(task *gtask.Task) (string, error)
	GetTasksFn    func() (*gtask.TasksData, error)
	GetTaskByIDFn func(id string) (*gtask.Task, error)
	UpdateTaskFn  func(id string, task *gtask.Task) error
	DeleteTaskFn  func(id string) error
}

func (s *Service) CreateTask(task *gtask.Task) (string, error) {
	return s.CreateTaskFn(task)
}

func (s *Service) GetTasks() (*gtask.TasksData, error) {
	return s.GetTasksFn()
}

func (s *Service) GetTaskByID(id string) (*gtask.Task, error) {
	return s.GetTaskByIDFn(id)
}

func (s *Service) UpdateTask(id string, task *gtask.Task) error {
	return s.UpdateTaskFn(id, task)
}

func (s *Service) DeleteTask(id string) error {
	return s.DeleteTaskFn(id)
}

// This ensures the Repository type implements the Repository interface via a compiler check,
// even if it is not used elsewhere.  You can read more on this pattern on the effective go site
// https://golang.org/doc/effective_go#blank_implements
var _ gtask.Repository = (*Repository)(nil)

type Repository struct {
	SaveFn     func(task *gtask.Task) (string, error)
	FindAllFn  func() (*gtask.TasksData, error)
	FindByIDFn func(id string) (*gtask.Task, error)
	UpdateFn   func(id string, task *gtask.Task) error
	DeleteFn   func(id string) error
}

func (r *Repository) Save(task *gtask.Task) (string, error) {
	return r.SaveFn(task)
}

func (r *Repository) FindAll() (*gtask.TasksData, error) {
	return r.FindAllFn()
}

func (r *Repository) FindByID(id string) (*gtask.Task, error) {
	return r.FindByIDFn(id)
}

func (r *Repository) Update(id string, task *gtask.Task) error {
	return r.UpdateFn(id, task)
}

func (r *Repository) Delete(id string) error {
	return r.DeleteFn(id)
}
