package file

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
)

func currentFilePath(t *testing.T) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Error("Can not read caller file path")
	}
	return filename
}

func currentDirectoryPath(t *testing.T) string {
	return filepath.Dir(currentFilePath(t))
}

func temporaryDirectoryPath(t *testing.T) string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}
	return filepath.Join("/tmp", "itreetest+"+uuid.String())
}
