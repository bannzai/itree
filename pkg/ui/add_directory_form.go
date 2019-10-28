package ui

import "github.com/rivo/tview"

type AddDirectoryForm struct {
	*tview.Form
}

func (form AddDirectoryForm) form() *tview.Form {
	return form.Form
}
func (form AddDirectoryForm) isDir() bool {
	return true
}
func (form AddDirectoryForm) makeFunction(path string) error {
	return makeDirectory(path)
}
func (AddDirectoryForm) name() string {
	return nameOfAddForm
}
func (form AddDirectoryForm) view() tview.Primitive {
	return form.Form
}
func (form AddDirectoryForm) decideButtonTitle() string {
	return "Add Directory"
}

func (window *Window) SwitchAddDirectoryForm(selectedNode *tview.TreeNode) {
	form := AddDirectoryForm{tview.NewForm()}
	window.switchAddFileForm(
		form,
		selectedNode,
	)
}
