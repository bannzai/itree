package ui

import (
	"github.com/rivo/tview"
)

type Page interface {
	name() string
	view() tview.Primitive
}

type Pages struct {
	*tview.Pages
}

func NewPages(root *MainView, pages ...Page) Pages {
	pagesView := tview.NewPages()
	pagesView.AddPage("main", root, true, true)
	for _, page := range pages {
		pagesView.AddPage(page.name(), page.view(), true, false)
	}
	return Pages{
		Pages: pagesView,
	}
}
