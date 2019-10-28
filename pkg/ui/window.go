package ui

import (
	"github.com/rivo/tview"
)

type switcher interface {
	SwitchRenameForm(node *tview.TreeNode)
	SwitchAddFileForm(*tview.TreeNode)
	SwitchAddDirectoryForm(*tview.TreeNode)
}

type Root interface {
	tview.Primitive
	AddAndSwitchToPage(name string, item tview.Primitive, resize bool) *tview.Pages
	RemovePage(name string) *tview.Pages
}

type Window struct {
	Root
}

func NewWindow(width, height int) *Window {
	window := &Window{}
	window.Root = NewPages(
		NewTree(window),
	)
	window.SetRect(0, 0, width, height)
	return window
}
