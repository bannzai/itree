package ui

import "github.com/rivo/tview"

const (
	inputWidth = 100
)

type Form struct {
	*tview.Form
}

func (window *Window) SwitchRenameForm(node *tview.TreeNode) Form {
	form := tview.NewForm()
	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	text := node.GetReference().(string)
	form.
		AddInputField("New Path", text, inputWidth, nil, func(s string) {
			text = s
		}).
		AddButton("Decide", func() {
			node.SetReference(text)
		}).
		AddButton("Cancel", func() {

		})
	return Form{
		form,
	}
}

func (Form) name() string {
	return nameOfForm
}

func (form Form) view() tview.Primitive {
	return form.Form
}
