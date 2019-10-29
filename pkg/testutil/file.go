package testutil

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/uuid"
)

func CurrentFilePath(t *testing.T) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Error("Can not read caller file path")
	}
	return filename
}

func CurrentDirectoryPath(t *testing.T) string {
	return filepath.Dir(CurrentFilePath(t))
}

func TemporaryDirectoryPath(t *testing.T) string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}
	return filepath.Join("/tmp", "itreetest+"+uuid.String())
}
