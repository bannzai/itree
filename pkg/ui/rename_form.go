package ui

import (
	"path/filepath"

	"github.com/bannzai/itree/pkg/file"
	"github.com/rivo/tview"
)

type RenameForm struct {
	*tview.Form
}

func (RenameForm) name() string {
	return nameOfRenameForm
}

func (form RenameForm) view() tview.Primitive {
	return form.Form
}

func (window *Window) SwitchRenameForm(node *tview.TreeNode) {
	const inputWidth = 100

	form := RenameForm{tview.NewForm()}

	closeForm := func() {
		window.Root.RemovePage(form.name())
	}

	nodeReference := extractNodeReference(node)
	fromPath := nodeReference.path
	directoryPath := filepath.Dir(fromPath)
	editedFileName := filepath.Base(fromPath)

	errorField := NewErrorField()

	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	form.
		AddInputField("New Path", editedFileName, inputWidth, nil, func(s string) {
			editedFileName = s
		}).
		AddButton("Decide", func() {
			editedPath := filepath.Join(directoryPath, editedFileName)
			if fromPath == editedPath {
				closeForm()
				return
			}

			if err := file.MoveFile(fromPath, editedPath); err != nil {
				errorField.SetText(err.Error())
				return
			}

			nodeReference.setPath(editedPath)
			node.SetReference(nodeReference)
			node.SetText(editedFileName)
			closeForm()
		}).
		AddButton("Cancel", func() {
			closeForm()
		}).
		SetCancelFunc(closeForm)

	grid := NewFormLayout(form, errorField)
	window.Root.AddAndSwitchToPage(form.name(), grid, true)

}
