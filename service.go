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

func (s *TaskService) CreateTask(task *Task) (string, error) {
	return s.repo.Save(task)
}

func (s *TaskService) GetTasks() (*TasksData, error) {
	return s.repo.FindAll()
}

func (s *TaskService) GetTaskByID(id string) (*Task, error) {
	return s.repo.FindByID(id)
}

func (s *TaskService) UpdateTask(id string, task *Task) error {
	return s.repo.Update(id, task)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
