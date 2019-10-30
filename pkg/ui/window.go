package ui

import (
	"github.com/gdamore/tcell"
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
	Root
	Transition
}

func NewWindow(width, height int) *Window {
	window := &Window{}
	window.Root = NewMainView(window)
	window.Transition = NewPages(
		window.Root,
	)
	window.SetRect(0, 0, width, height)
	return window
}

// Confirm for tview.Primitive
func (window Window) Draw(screen tcell.Screen) {
	window.Transition.Draw(screen)
}
func (window Window) GetRect() (x int, y int, width int, height int) {
	return window.Transition.GetRect()
}
func (window Window) SetRect(x, y, width, height int) {
	window.Transition.SetRect(x, y, width, height)
}
func (window Window) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return window.Transition.InputHandler()
}
func (window Window) Focus(delegate func(p tview.Primitive)) {
	window.Transition.Focus(delegate)
}
func (window Window) Blur() {
	window.Transition.Blur()
}
func (window Window) GetFocusable() tview.Focusable {
	return window.Transition.GetFocusable()
}
