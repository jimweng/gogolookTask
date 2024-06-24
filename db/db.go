package db

import (
	"io/fs"
	"os"

	task "github.com/jimweng/gogolookTask"
)

const rootPath = "./task/"

type FileSystem interface {
	WriteFile(filename string, data []byte, perm fs.FileMode) error
	RemoveAll(path string) error
	MkdirAll(path string, perm fs.FileMode) error
	Remove(path string) error
	Rename(oldpath, newpath string) error
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

func (OSFileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (OSFileSystem) Remove(path string) error {
	return os.Remove(path)
}

func (OSFileSystem) Rename(oldpath, newpath string) error {
	return os.Rename(oldpath, newpath)
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

func (r *Repository) Save(task *task.Task) (string, error) {
	return "", nil
}
func (r *Repository) FindAll() ([]*task.Task, error) {
	return nil, nil
}
func (r *Repository) FindByID(id string) (*task.Task, error) {
	return nil, nil
}
func (r *Repository) Update(task *task.Task) error {
	return nil
}
func (r *Repository) Delete(id string) error {
	return nil
}
