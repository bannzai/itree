package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"
	"time"
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
	return filepath.Join("/tmp", fmt.Sprintf("itreetest+%s_%d", t.Name(), time.Now().Unix()))
}

func pathForCreateTemporaryFile(t *testing.T) string {
	file, err := ioutil.TempFile("/tmp", fmt.Sprintf("itreetest+%s_%d", t.Name(), time.Now().Unix()))
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
