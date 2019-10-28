package ui

import (
	"os"
	"os/exec"
	"path/filepath"

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

func moveFile(from, to string) error {
	mvCmd := exec.Command("mv", from, to)
	return mvCmd.Run()
}

func makeFile(path string) error {
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
	return tview.NewTreeNode(fileName).
		SetReference(
			newNodeReference(
				filepath.Join(parentPath, fileName),
				isDir,
				parent,
			),
		).
		SetSelectable(true)
}
