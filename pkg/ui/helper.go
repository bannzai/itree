package ui

import (
	"os/exec"
	"path/filepath"
)

func absolutePath(node nodeReference) string {
	relativePath := node.path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	return absolutePath
}

func moveFile(from, to string) error {
	mvCmd := exec.Command("mv", from, to)
	return mvCmd.Run()
}
