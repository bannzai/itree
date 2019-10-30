package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type transition interface {
	tview.Primitive
	AddAndSwitchToPage(name string, item tview.Primitive, resize bool) *tview.Pages
	RemovePage(name string) *tview.Pages
}

type Window struct {
	*Root
	transition
}

func NewWindow(width, height int) *Window {
	window := &Window{}
	window.Root = NewRoot(window)
	window.transition = NewPages(
		window.Root,
	)
	window.SetRect(0, 0, width, height)
	return window
}

// Confirm for tview.Primitive
func (window Window) Draw(screen tcell.Screen) {
	window.transition.Draw(screen)
}
func (window Window) GetRect() (x int, y int, width int, height int) {
	return window.transition.GetRect()
}
func (window Window) SetRect(x, y, width, height int) {
	window.transition.SetRect(x, y, width, height)
}
func (window Window) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return window.transition.InputHandler()
}
func (window Window) Focus(delegate func(p tview.Primitive)) {
	window.transition.Focus(delegate)
}
func (window Window) Blur() {
	window.transition.Blur()
}
func (window Window) GetFocusable() tview.Focusable {
	return window.transition.GetFocusable()
}
