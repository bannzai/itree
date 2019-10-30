package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Page interface {
	name() string
	view() tview.Primitive
}

type Pages struct {
	*tview.Pages
	Root
}

func NewPages(root Root, pages ...Page) Pages {
	pagesView := tview.NewPages()
	pagesView.AddPage("main", root, true, true)
	for _, page := range pages {
		pagesView.AddPage(page.name(), page.view(), true, false)
	}
	return Pages{
		Pages: pagesView,
		Root:  root,
	}
}
func (view Pages) Draw(screen tcell.Screen) {
	view.Pages.Draw(screen)
}

func (view Pages) GetRect() (x int, y int, width int, height int) {
	return view.Pages.GetRect()
}

func (view Pages) SetRect(x, y, width, height int) {
	view.Pages.SetRect(x, y, width, height)
}
func (view Pages) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return view.Pages.InputHandler()
}
func (view Pages) Focus(delegate func(p tview.Primitive)) {
	view.Pages.Focus(delegate)
}
func (view Pages) Blur() {
	view.Pages.Blur()
}
func (view Pages) GetFocusable() tview.Focusable {
	return view.Pages.GetFocusable()
}
