package ui

import (
	"fmt"
	"path/filepath"

	"github.com/rivo/tview"
)

func (window *Window) SwitchRenameForm(node *tview.TreeNode) {
	const inputWidth = 100

	form := Form{tview.NewForm()}

	closeForm := func() {
		window.Root.RemovePage(form.name())
	}

	nodeReference := node.GetReference().(nodeReference)
	fromPath := nodeReference.path
	directoryPath := filepath.Dir(fromPath)
	editedFileName := filepath.Base(fromPath)

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

			if err := moveFile(fromPath, editedPath); err != nil {
				panic(fmt.Sprintf("err %v, from: %s, editedFileName: %s, editedPath: %s", err, fromPath, editedFileName, editedPath))
			} else {
				// TODO: show error dialog
				nodeReference.setPath(editedPath)
				node.SetReference(nodeReference)
				node.SetText(editedFileName)
				closeForm()
			}
		}).
		AddButton("Cancel", func() {
			closeForm()
		})

	window.Root.AddAndSwitchToPage(form.name(), form.view(), true)
}

func (Form) name() string {
	return nameOfForm
}

func (form Form) view() tview.Primitive {
	return form.Form
}
