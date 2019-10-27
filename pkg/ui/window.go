package ui

import "github.com/rivo/tview"

type Root interface {
	tview.Primitive
	AddAndSwitchToPage(name string, item tview.Primitive, resize bool) *tview.Pages
	RemovePage(name string) *tview.Pages
}

type Window struct {
	tview.Primitive
	root Root
}

func NewWindow() Window {
	return Window{
		root: NewPages(
			NewTree(),
		),
	}
}
