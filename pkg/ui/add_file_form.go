package ui

import (
	"fmt"
	"path/filepath"

	"github.com/rivo/tview"
)

func (window *Window) SwitchAddFileForm(selectedNode *tview.TreeNode, makeFileOrDirectory func(string) error) {
	const inputWidth = 100

	form := Form{tview.NewForm()}

	closeForm := func() {
		window.Root.RemovePage(form.name())
	}

	var directoryNode *tview.TreeNode
	selectedNodeReference := extractNodeReference(selectedNode)
	if selectedNodeReference.isDir {
		directoryNode = selectedNode
	} else {
		directoryNode = selectedNodeReference.parentNode
	}
	directoryNodeReference := extractNodeReference(directoryNode)
	directoryPath := directoryNodeReference.path
	editedFileName := ""

	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	form.
		AddInputField("Add File", editedFileName, inputWidth, nil, func(s string) {
			editedFileName = s
		}).
		AddButton("Decide", func() {
			path := filepath.Join(directoryPath, editedFileName)
			if err := makeFileOrDirectory(path); err != nil {
				panic(fmt.Sprintf("err %v, directoryPath: %v, editedFileName: %s", err, directoryPath, editedFileName))
			} else {
				// TODO: show error dialog
				directoryNode.AddChild(createTreeNode(editedFileName, false, directoryNode))
				closeForm()
			}
		}).
		AddButton("Cancel", func() {
			closeForm()
		})

	window.Root.AddAndSwitchToPage(form.name(), form.view(), true)
}
