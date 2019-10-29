package ui

import (
	"path/filepath"

	"github.com/gdamore/tcell"
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
