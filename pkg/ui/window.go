package ui

import (
	"github.com/rivo/tview"
)

type Root interface {
	tview.Primitive
	ShowFileInfo(path string)
}

type Transition interface {
	tview.Primitive
	AddAndSwitchToPage(name string, item tview.Primitive, resize bool) *tview.Pages
	RemovePage(name string) *tview.Pages
}

type Window struct {
	Transition
}

func NewWindow(width, height int) *Window {
	window := &Window{}
	window.Transition = NewPages(
		NewMainView(window),
	)
	window.SetRect(0, 0, width, height)
	return window
}
