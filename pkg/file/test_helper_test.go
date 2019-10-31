package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
	"time"

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
	return filepath.Join("/tmp", fmt.Sprintf("itreetest+%s-%s_%d", t.Name(), uuid.String(), time.Now().Unix()))
}

func pathForCreatedTemporaryFile(t *testing.T) string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}
	file, err := ioutil.TempFile("/tmp", fmt.Sprintf("itreetest+%s-%s_%d", t.Name(), uuid.String(), time.Now().Unix()))
	if err != nil {
		t.Error(err)
	}
	if err := file.Close(); err != nil {
		t.Error(err)
	}
	return file.Name()
}

func temporaryFilePath(t *testing.T) string {
	return temporaryDirectoryPath(t)
}
