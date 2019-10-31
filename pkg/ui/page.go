package ui

import (
	"github.com/rivo/tview"
)

type page interface {
	name() string
	view() tview.Primitive
}

type pages struct {
	*tview.Pages
}

func newPages(root *root, pageViews ...page) pages {
	pagesView := tview.NewPages()
	pagesView.AddPage("main", root, true, true)
	for _, page := range pageViews {
		pagesView.AddPage(page.name(), page.view(), true, false)
	}
	return pages{
		Pages: pagesView,
	}
}
