package task

type Task struct {
	Name   string `json:"name,omitempty"`
	Status Status `json:"status,omitempty"`
}

type TasksData struct {
	Tasks map[string]Task `json:"tasks,omitempty"`
}

type Service interface {
	CreateTask(id string, task *Task) (string, error)
	GetTasks() (*TasksData, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(id string, task *Task) error
	DeleteTask(id string) error
}

type Repository interface {
	Save(id string, task *Task) (string, error)
	FindAll() (*TasksData, error)
	FindByID(id string) (*Task, error)
	Update(id string, task *Task) error
	Delete(id string) error
}
