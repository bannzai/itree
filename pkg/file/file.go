package file

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func MoveFile(from, to string) error {
	mvCmd := exec.Command("mv", from, to)
	return mvCmd.Run()
}

func MakeFile(path string) error {
	if FileExists(path) {
		return fmt.Errorf("%s is already exists", path)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return errors.Wrapf(err, "New file error. path for  %s", path)
	}
	defer file.Close()
	return nil
}
