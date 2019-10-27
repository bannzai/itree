package ui

import "github.com/rivo/tview"

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

	form.
		SetBorder(true).
		SetTitleAlign(tview.AlignLeft).
		SetTitle("Rename")
	text := node.GetReference().(nodeReference).path
	form.
		AddInputField("New Path", text, inputWidth, nil, func(s string) {
			text = s
		}).
		AddButton("Decide", func() {
			from := node.GetReference().(nodeReference).path
			node.SetReference(text)
			moveFile(from, text)
			closeForm()
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
