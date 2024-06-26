package mock

import (
	"io/fs"
	"os"
)

type FileSystem struct {
	WriteFileFn func(filename string, data []byte, perm fs.FileMode) error
	CreateFn    func(name string) (*os.File, error)
	ReadFileFn  func(name string) ([]byte, error)
	OepnFn      func(name string) (*os.File, error)
}

func (fs FileSystem) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return fs.WriteFileFn(filename, data, perm)
}

func (fs FileSystem) Create(name string) (*os.File, error) {
	return fs.CreateFn(name)
}

func (fs FileSystem) ReadFile(name string) ([]byte, error) {
	return fs.ReadFileFn(name)
}

func (fs FileSystem) Open(name string) (*os.File, error) {
	return fs.OepnFn(name)
}
