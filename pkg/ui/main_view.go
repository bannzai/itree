package ui

import "github.com/rivo/tview"

type MainView struct {
	*tview.Grid
	FileInfo
}

func NewMainView(window *Window) MainView {
	tree := NewTree(window)
	view := MainView{
		Grid: tview.NewGrid().
			SetRows(0, 1).
			SetColumns(0).
			AddItem(tree, 0, 0, 1, 1, 0, 0, true),
	}
	return view
}

func (view MainView) ShowFileInfo(path string) {
	if view.FileInfo.List != nil {
		view.RemoveItem(view.FileInfo.List)
	}

	fileInfo := NewFileInfo(path)
	view.AddItem(fileInfo.List, 0, 1, 1, 1, 0, 0, false)
	view.FileInfo = fileInfo
}
