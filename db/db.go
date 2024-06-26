package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/google/uuid"
	gtask "github.com/jimweng/gogolookTask"
)

type FileSystem interface {
	WriteFile(name string, data []byte, perm fs.FileMode) error
	Create(name string) (*os.File, error)
	ReadFile(name string) ([]byte, error)
	Open(name string) (*os.File, error)
}


type OSFileSystem struct{}

func (OSFileSystem) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (OSFileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (OSFileSystem) Open(name string) (*os.File, error) {
	return os.Open(name)
}

type Repository struct {
	filePath   string
	fileSystem FileSystem
}

func NewFileSystem() *OSFileSystem {
	return &OSFileSystem{}
}

func NewRepository(filePath string, fs FileSystem) *Repository {
	return &Repository{
		filePath:   filePath,
		fileSystem: fs,
	}
}

func (r *Repository) Save(task *gtask.Task) (string, error) {
	taskData, err := r.loadTasks()
	if err != nil {
		return "", errors.New("failed to read file")
	}

	for _, v := range taskData.Tasks {
		if v.Name == task.Name {
			return "", errors.New("the taskname is existed")
		}
	}

	id := uuid.New().String()
	if taskData.Tasks == nil {
		taskData.Tasks = make(map[string]*gtask.Task)
	}

	taskData.Tasks[id] = task

	if err := r.saveTasks(taskData); err != nil {
		return "", errors.New("failed to save file")
	}
	return id, nil
}

func (r *Repository) FindAll() (*gtask.TasksData, error) {
	return r.loadTasks()
}

func (r *Repository) FindByID(id string) (*gtask.Task, error) {
	taskData, err := r.loadTasks()
	if err != nil {
		return nil, errors.New("failed to read file")
	}

	task, exist := taskData.Tasks[id]
	if !exist {
		return nil, errors.New("not found")
	}
	return task, nil
}

func (r *Repository) Update(id string, task *gtask.Task) error {
	taskData, err := r.loadTasks()
	if err != nil {
		return errors.New("failed to read file")
	}

	if _, exist := taskData.Tasks[id]; !exist {
		return errors.New("id not found")
	}

	taskData.Tasks[id] = task

	return r.saveTasks(taskData)
}

func (r *Repository) Delete(id string) error {
	taskData, err := r.loadTasks()
	if err != nil {
		return errors.New("failed to read file")
	}

	delete(taskData.Tasks, id)

	return r.saveTasks(taskData)
}

func (r *Repository) loadTasks() (*gtask.TasksData, error) {
	file, err := r.fileSystem.Open(r.filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			f, err := r.fileSystem.Create(r.filePath)
			if err != nil {
				return nil, errors.New("failed to create file")
			}
			defer f.Close()

			jsonData, err := json.Marshal(&gtask.TasksData{
				Tasks: map[string]*gtask.Task{},
			})
			if err != nil {
				return nil, errors.New("failed to marshal filesystem")
			}

			_, err = f.Write(jsonData)
			if err != nil {
				return nil, errors.New("failed to write to file")
			}

			return r.loadTasks()
		}
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := r.fileSystem.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var tasksData gtask.TasksData
	err = json.Unmarshal(byteValue, &tasksData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &tasksData, nil
}

func (r *Repository) saveTasks(taskData *gtask.TasksData) error {
	bytes, err := json.Marshal(taskData)
	if err != nil {
		return errors.New("failed to marshal task data")
	}

	if err = r.fileSystem.WriteFile(r.filePath, bytes, fs.ModePerm); err != nil {
		return errors.New("failed to write file")
	}
	return nil
}
