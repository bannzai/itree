package ui

import (
	"github.com/rivo/tview"
)

type addFileForm interface {
	Page
	form() *tview.Form
	isDir() bool
	makeFunction(path string) error
}

type AddFileForm struct {
	*tview.Form
}

func (form AddFileForm) form() *tview.Form {
	return form.Form
}
func (form AddFileForm) isDir() bool {
	return false
}
func (form AddFileForm) makeFunction(path string) error {
	return makeFile(path)
}
func (AddFileForm) name() string {
	return nameOfAddForm
}
func (form AddFileForm) view() tview.Primitive {
	return form.Form
}

func (window *Window) SwitchAddFileForm(selectedNode *tview.TreeNode) {

	form := AddFileForm{tview.NewForm()}
	window.switchAddFileForm(
		form,
		selectedNode,
	)
}
