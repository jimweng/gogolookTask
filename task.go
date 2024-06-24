package task

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

type Service interface {
	CreateTask(task *Task) (string, error)
	GetTasks() ([]*Task, error)
	GetTaskByID(id string) (*Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
}

type Repository interface {
	Save(task *Task) (string, error)
	FindAll() ([]*Task, error)
	FindByID(id string) (*Task, error)
	Update(task *Task) error
	Delete(id string) error
}
