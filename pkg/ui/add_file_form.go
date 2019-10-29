package ui

import (
	"github.com/bannzai/itree/pkg/file"
	"github.com/rivo/tview"
)

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
	return file.MakeFile(path)
}
func (AddFileForm) name() string {
	return nameOfAddForm
}
func (form AddFileForm) view() tview.Primitive {
	return form.Form
}
func (form AddFileForm) decideButtonTitle() string {
	return "Add File"
}

func (window *Window) SwitchAddFileForm(selectedNode *tview.TreeNode) {

	form := AddFileForm{tview.NewForm()}
	window.switchAddFileForm(
		form,
		selectedNode,
	)
}
