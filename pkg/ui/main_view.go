package ui

import (
	"github.com/rivo/tview"
)

type MainView struct {
	*tview.Grid
	FileInfo
}

func NewMainView(window *Window) *MainView {
	tree := NewTree(window)
	view := &MainView{
		Grid: tview.NewGrid().
			SetRows(0).
			SetColumns(30, 0).
			AddItem(tree, 0, 0, 1, 1, 0, 0, true),
	}
	return view
}

func (view *MainView) ShowFileInfo(path string) {
	if view.FileInfo.View != nil {
		view.RemoveItem(view.FileInfo.View)
	}

	fileInfo := NewFileInfo(path)
	view.AddItem(fileInfo.View, 0, 1, 1, 1, 0, 0, false)
	view.FileInfo = fileInfo
}
