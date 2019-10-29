package ui

import (
	"os"
	"os/exec"
)

type Editor struct{}

func NewEditor() Editor {
	return Editor{}
}

func (Editor) Launch(path string) {
	SharedConfig.Application.Suspend(func() {
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "vim"
		}
		cmd := exec.Command(editor, path)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Run()
	})
}
