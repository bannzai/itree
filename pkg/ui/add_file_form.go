package ui

import (
	"github.com/bannzai/itree/pkg/file"
	"github.com/rivo/tview"
)

type addFileForm struct {
	*tview.Form
}

func (form addFileForm) form() *tview.Form {
	return form.Form
}
func (form addFileForm) isDir() bool {
	return false
}
func (form addFileForm) makeFunction(path string) error {
	return file.MakeFile(path)
}
func (addFileForm) name() string {
	return nameOfAddForm
}
func (form addFileForm) view() tview.Primitive {
	return form.Form
}
func (form addFileForm) decideButtonTitle() string {
	return "Add File"
}

func (window *Window) SwitchAddFileForm(selectedNode *tview.TreeNode) {
	form := addFileForm{tview.NewForm()}
	window.switchAddFileForm(
		form,
		selectedNode,
	)
}
