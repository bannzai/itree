package ui

import "github.com/rivo/tview"

type Page interface {
	name() string
	view() tview.Primitive
}

type Pages struct {
	*tview.Pages
}

func NewPages(pages ...Page) Pages {
	pagesView := tview.NewPages()
	for i, page := range pages {
		pagesView.AddPage(page.name(), page.view(), false, i == 0)
	}
	return Pages{
		pagesView,
	}
}
