package ui

import (
	"github.com/bannzai/itree/pkg/file"
	"github.com/rivo/tview"
)

type addDirectoryForm struct {
	*tview.Form
}

func (form addDirectoryForm) form() *tview.Form {
	return form.Form
}
func (form addDirectoryForm) isDir() bool {
	return true
}
func (form addDirectoryForm) makeFunction(path string) error {
	return file.MakeDirectory(path)
}
func (addDirectoryForm) name() string {
	return nameOfAddForm
}
func (form addDirectoryForm) view() tview.Primitive {
	return form.Form
}
func (form addDirectoryForm) decideButtonTitle() string {
	return "Add Directory"
}

func (window *Window) SwitchAddDirectoryForm(selectedNode *tview.TreeNode) {
	form := addDirectoryForm{tview.NewForm()}
	window.switchAddFileForm(
		form,
		selectedNode,
	)
}
