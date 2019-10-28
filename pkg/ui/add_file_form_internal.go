package ui

import (
	"fmt"
	"path/filepath"

	"github.com/rivo/tview"
)

func (window *Window) switchAddFileForm(formView addFileForm, selectedNode *tview.TreeNode) {
	const inputWidth = 100

	form := formView.form()

	closeForm := func() {
		window.Root.RemovePage(formView.name())
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
			if err := formView.makeFunction(path); err != nil {
				panic(fmt.Sprintf("err %v, directoryPath: %v, editedFileName: %s", err, directoryPath, editedFileName))
			} else {
				// TODO: show error dialog
				directoryNode.AddChild(createTreeNode(editedFileName, formView.isDir(), directoryNode))
				closeForm()
			}
		}).
		AddButton("Cancel", func() {
			closeForm()
		})

	window.Root.AddAndSwitchToPage(formView.name(), formView.view(), true)
}