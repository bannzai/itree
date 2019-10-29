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

	errorField := NewErrorField()
	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	form.
		AddInputField(formView.decideButtonTitle(), editedFileName, inputWidth, nil, func(s string) {
			editedFileName = s
		}).
		AddButton("Decide", func() {
			path := filepath.Join(directoryPath, editedFileName)
			if err := formView.makeFunction(path); err != nil {
				errorField.SetText(err.Error())
				// panic(fmt.Sprintf("err %v, directoryPath: %v, editedFileName: %s", err, directoryPath, editedFileName))
			} else {
				// TODO: show error dialog
				directoryNode.AddChild(createTreeNode(editedFileName, formView.isDir(), directoryNode))
				closeForm()
			}
		}).
		AddButton("Cancel", func() {
			closeForm()
		}).
		SetCancelFunc(closeForm)

	grid := tview.NewGrid().
		SetRows(0, 1).
		SetColumns(0).
		SetBorders(true).
		AddItem(form, 0, 0, 1, 1, 0, 0, true).
		AddItem(errorField, 1, 0, 1, 1, 0, 30, false)

	window.Root.AddAndSwitchToPage(formView.name(), grid, true)
}
