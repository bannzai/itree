package ui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gdamore/tcell"
	"github.com/pkg/errors"
	"github.com/rivo/tview"
)

func absolutePath(node nodeReference) string {
	relativePath := node.path
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
	}
	return absolutePath
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func moveFile(from, to string) error {
	mvCmd := exec.Command("mv", from, to)
	return mvCmd.Run()
}

func makeFile(path string) error {
	if fileExists(path) {
		return fmt.Errorf("%s is already exists", path)
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return errors.Wrapf(err, "New file error. path for  %s", path)
	}
	defer file.Close()
	return nil
}

func makeDirectory(path string) error {
	return os.Mkdir(path, 0777)
}

func extractNodeReference(node *tview.TreeNode) *nodeReference {
	return node.GetReference().(*nodeReference)
}

func createTreeNode(fileName string, isDir bool, parent *tview.TreeNode) *tview.TreeNode {
	var parentPath string
	parentReference, ok := parent.GetReference().(*nodeReference)
	if ok {
		parentPath = parentReference.path
	} else {
		parentPath = "./"
	}

	var color tcell.Color
	if isDir {
		color = tcell.ColorGreen
	} else {
		color = tview.Styles.PrimaryTextColor
	}

	return tview.NewTreeNode(fileName).
		SetReference(
			newNodeReference(
				filepath.Join(parentPath, fileName),
				isDir,
				parent,
			),
		).
		SetSelectable(true).
		SetColor(color)

}
