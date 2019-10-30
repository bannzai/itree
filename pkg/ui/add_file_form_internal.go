package ui

import (
	"path/filepath"

	"github.com/rivo/tview"
)

type addFileForm interface {
	Page
	form() *tview.Form
	isDir() bool
	makeFunction(path string) error
	decideButtonTitle() string
}

func (window *Window) switchAddFileForm(formView addFileForm, selectedNode *tview.TreeNode) {
	const inputWidth = 100

	form := formView.form()

	closeForm := func() {
		window.transition.RemovePage(formView.name())
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

	errorField := NewErrorField()
	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Add")
	form.
		AddInputField(formView.decideButtonTitle(), editedFileName, inputWidth, nil, func(s string) {
			editedFileName = s
		}).
		AddButton("Decide", func() {
			path := filepath.Join(directoryPath, editedFileName)
			if err := formView.makeFunction(path); err != nil {
				errorField.SetText(err.Error())
				return
			}

			directoryNode.AddChild(createTreeNode(editedFileName, formView.isDir(), directoryNode))
			closeForm()
		}).
		AddButton("Cancel", func() {
			closeForm()
		}).
		SetCancelFunc(closeForm)

	grid := NewFormLayout(form, errorField)
	window.transition.AddAndSwitchToPage(formView.name(), grid, true)
}
