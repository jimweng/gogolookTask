package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"

	task "github.com/jimweng/gogolookTask"
)

type FileSystem interface {
	WriteFile(filename string, data []byte, perm fs.FileMode) error
	RemoveAll(path string) error
	Remove(path string) error
	Create(name string) (*os.File, error)
	ReadFile(name string) ([]byte, error)
}

type OSFileSystem struct{}

func (OSFileSystem) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (OSFileSystem) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (OSFileSystem) Remove(path string) error {
	return os.Remove(path)
}

func (OSFileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}

func (OSFileSystem) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
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


// TODO: duplicate check
func (r *Repository) Save(id string, task *task.Task) (string, error) {
	taskData, err := loadTasks(r.filePath)
	if err != nil {
		return "", errors.New("failed to read file")
	}

	if _, exists := taskData.Tasks[id]; exists {
		return "", errors.New("the task already exists")
	}
	return id, nil
}

func (r *Repository) FindAll() (*task.TasksData, error) {
	return loadTasks(r.filePath)
}

func (r *Repository) FindByID(id string) (*task.Task, error) {
	taskData, err := loadTasks(r.filePath)
	if err != nil {
		return nil, errors.New("failed to read file")
	}

	task, _ := taskData.Tasks[id]
	return &task, nil
}

func (r *Repository) Update(id string, task *task.Task) error {
	taskData, err := loadTasks(r.filePath)
	if err != nil {
		return errors.New("failed to read file")
	}

	taskData.Tasks[id] = *task

	return r.saveTasks(taskData)
}

func (r *Repository) Delete(id string) error {
	taskData, err := loadTasks(r.filePath)
	if err != nil {
		return errors.New("failed to read file")
	}

	delete(taskData.Tasks, id)

	return r.saveTasks(taskData)
}

func loadTasks(filePath string) (*task.TasksData, error) {
	var tasksData task.TasksData
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	err = json.Unmarshal(byteValue, &tasksData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &tasksData, nil
}

func (r *Repository) saveTasks(taskData *task.TasksData) error {
	bytes, err := json.Marshal(taskData)
	if err != nil {
		return errors.New("failed to marshal task data")
	}

	if err = r.fileSystem.WriteFile(r.filePath, bytes, fs.ModePerm); err != nil {
		return errors.New("failed to write file")
	}
	return nil
}