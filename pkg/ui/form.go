package ui

import (
	"path/filepath"

	"github.com/rivo/tview"
)

const (
	inputWidth = 100
)

type Form struct {
	*tview.Form
}

func (window *Window) SwitchRenameForm(node *tview.TreeNode) {
	form := Form{tview.NewForm()}

	closeForm := func() {
		window.Root.RemovePage(form.name())
	}

	nodeReference := node.GetReference().(nodeReference)
	fromPath := nodeReference.path
	directoryPath := filepath.Dir(fromPath)
	editedPath := filepath.Base(fromPath)

	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	form.
		AddInputField("New Path", editedPath, inputWidth, nil, func(s string) {
			editedPath = s
		}).
		AddButton("Decide", func() {
			if fromPath == editedPath {
				closeForm()
				return
			}

			if err := moveFile(fromPath, editedPath); err == nil {
				// TODO: show error dialog
				nodeReference.setPath(filepath.Join(directoryPath, editedPath))
				node.SetReference(nodeReference)
				node.SetText(editedPath)
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
