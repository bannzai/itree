package ui

import (
	"path/filepath"

	"github.com/bannzai/itree/pkg/file"
	"github.com/rivo/tview"
)

type renameForm struct {
	*tview.Form
}

func (renameForm) name() string {
	return nameOfRenameForm
}

func (form renameForm) view() tview.Primitive {
	return form.Form
}

func (window *Window) SwitchRenameForm(node *tview.TreeNode) {
	const inputWidth = 100

	form := renameForm{tview.NewForm()}

	closeForm := func() {
		window.transition.RemovePage(form.name())
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

			if nodeReference.isDir {
				for _, child := range node.GetChildren() {
					childReference := extractNodeReference(child)
					basename := filepath.Base(childReference.path)
					childPath := filepath.Join(editedPath, basename)
					childReference.path = childPath
					child.SetReference(childReference)
				}
			}

			nodeReference.path = editedPath
			node.SetReference(nodeReference)
			node.SetText(editedFileName)
			closeForm()
		}).
		AddButton("Cancel", func() {
			closeForm()
		}).
		SetCancelFunc(closeForm)

	grid := NewFormLayout(form, errorField)
	window.transition.AddAndSwitchToPage(form.name(), grid, true)

}
