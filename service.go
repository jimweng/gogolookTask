package task

type TaskService struct {
	repo Repository
}

func NewService(repo Repository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func Must(service *Service, err error) *Service {
	if err != nil {
		panic(err)
	}
	return service
}

func (s *TaskService) CreateTask(task *Task) error {
	return nil
}

func (s *TaskService) GetTasks() ([]*Task, error) {
	return nil, nil
}

func (s *TaskService) GetTaskByID(id string) (*Task, error) {
	return nil, nil
}

func (s *TaskService) UpdateTask(task *Task) error {
	return nil
}

func (s *TaskService) DeleteTask(id string) error {
	return nil
}
