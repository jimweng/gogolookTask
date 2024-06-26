package db_test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"testing"

	gtask "github.com/jimweng/gogolookTask"
	gtaskdb "github.com/jimweng/gogolookTask/db"
	"github.com/jimweng/gogolookTask/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testFileName = "testFileName"
	testName     = "testName"
	testStatus   = 0
	testID       = "testId"
	errMsg       = "some error"
)

func Test_NewFileSystem(t *testing.T) {
	t.Parallel()
	fs := gtaskdb.NewFileSystem()
	assert.NotNil(t, fs)
}

func Test_NewRepository(t *testing.T) {
	t.Parallel()
	fs := mock.FileSystem{}
	repo := gtaskdb.NewRepository(testFileName, fs)
	assert.NotNil(t, repo)
}

//nolint:paralleltest,tparallel
func Test_Save(t *testing.T) {
	t.Parallel()
	t.Run("successfully save file", func(t *testing.T) {
		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.Save(&gtask.Task{})
		require.NoError(t, err)
	})

	t.Run("failed to read file", func(t *testing.T) {
		mfs := mock.FileSystem{
			OepnFn: func(_ string) (*os.File, error) {
				return nil, fmt.Errorf(errMsg)
			},
		}

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.Save(&gtask.Task{})
		require.Error(t, err, errMsg)
	})

	t.Run("the taskname is existed", func(t *testing.T) {
		mfs := mockFileSystem()

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.Save(&gtask.Task{
			Name:   testName,
			Status: testStatus,
		})
		require.Error(t, err, "the taskname is existed")
	})

	t.Run("failed to save tasks", func(t *testing.T) {
		mfs := mock.FileSystem{
			OepnFn: func(_ string) (*os.File, error) {
				return nil, fmt.Errorf(errMsg)
			},
		}

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.Save(&gtask.Task{})
		require.Error(t, err, errMsg)
	})

	t.Run("the taskname is existed", func(t *testing.T) {
		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.Save(&gtask.Task{
			Name:   testName,
			Status: testStatus,
		})
		require.Error(t, err, "the taskname is existed")
	})
}

//nolint:paralleltest
func Test_FindAll(t *testing.T) {
	// Test loadTasks
	t.Run("successfully find all", func(t *testing.T) {
		mfs := mockFileSystem()

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindAll()
		require.NoError(t, err)
	})

	t.Run("failed to create file", func(t *testing.T) {
		mfs := mock.FileSystem{
			OepnFn: func(_ string) (*os.File, error) {
				return nil, os.ErrNotExist
			},
			CreateFn: func(_ string) (*os.File, error) {
				return nil, fmt.Errorf(errMsg)
			},
		}

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindAll()
		require.Error(t, err, errMsg)
	})

	t.Run("failed to write to file", func(t *testing.T) {
		mfs := mock.FileSystem{
			OepnFn: func(_ string) (*os.File, error) {
				return nil, os.ErrNotExist
			},
			CreateFn: func(_ string) (*os.File, error) {
				return &os.File{}, nil
			},
		}

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindAll()
		require.Error(t, err, "failed to write to file")
	})

	t.Run("failed to read file", func(t *testing.T) {
		mfs := mockFileSystem()
		mfs.ReadFileFn = func(_ string) ([]byte, error) {
			return nil, fmt.Errorf(errMsg)
		}

		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindAll()
		require.NotEmpty(t, err)
	})
}

//nolint:tparallel,paralleltest
func Test_FindByID(t *testing.T) {
	t.Parallel()
	t.Run("successfully find by id", func(t *testing.T) {
		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindByID(testID)
		require.NoError(t, err)
	})

	t.Run("failed to read file", func(t *testing.T) {
		mfs := mockFileSystem()
		mfs.ReadFileFn = func(_ string) ([]byte, error) {
			return nil, fmt.Errorf(errMsg)
		}
		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindByID(testID)
		require.Error(t, err, "faile to read file")
	})

	t.Run("id not found", func(t *testing.T) {
		wrongID := "wrong" + testID

		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		_, err := repo.FindByID(wrongID)
		require.Error(t, err, "id not found")
	})
}

//nolint:tparallel,paralleltest
func Test_Update(t *testing.T) {
	t.Parallel()
	t.Run("successfully update", func(t *testing.T) {
		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		err := repo.Update(testID, &gtask.Task{})
		require.NoError(t, err)
	})

	t.Run("failed to read file", func(t *testing.T) {
		mfs := mockFileSystem()
		mfs.ReadFileFn = func(_ string) ([]byte, error) {
			return nil, fmt.Errorf(errMsg)
		}
		repo := gtaskdb.NewRepository(testFileName, mfs)
		err := repo.Update(testID, &gtask.Task{})
		require.Error(t, err, "faile to read file")
	})

	t.Run("id not found", func(t *testing.T) {
		wrongID := "wrong" + testID

		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		err := repo.Update(wrongID, &gtask.Task{})
		require.Error(t, err, "id not found")
	})
}

//nolint:tparallel,paralleltest
func Test_Delete(t *testing.T) {
	t.Parallel()
	t.Run("successfully delete the task", func(t *testing.T) {
		mfs := mockFileSystem()
		repo := gtaskdb.NewRepository(testFileName, mfs)
		err := repo.Delete(testID)
		require.NoError(t, err)
	})

	t.Run("faile to read file", func(t *testing.T) {
		mfs := mockFileSystem()
		mfs.ReadFileFn = func(_ string) ([]byte, error) {
			return nil, fmt.Errorf(errMsg)
		}
		repo := gtaskdb.NewRepository(testFileName, mfs)
		err := repo.Delete(testID)
		require.Error(t, err, "faile to read file")
	})
}

func mockFileSystem() mock.FileSystem {
	return mock.FileSystem{
		OepnFn: func(_ string) (*os.File, error) {
			return nil, nil
		},
		CreateFn: func(_ string) (*os.File, error) {
			return nil, nil
		},
		ReadFileFn: func(_ string) ([]byte, error) {
			mtask := gtask.TasksData{
				Tasks: map[string]*gtask.Task{
					testID: {
						Name:   testName,
						Status: testStatus,
					},
				},
			}
			bytemTask, _ := json.Marshal(mtask)
			return bytemTask, nil
		},
		WriteFileFn: func(_ string, _ []byte, _ fs.FileMode) error {
			return nil
		},
	}
}
