package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type root struct {
	*tview.Grid
	fileInfo
	feedback
	search
	*tree
}

func newRoot(window *Window) *root {
	tree := newTree(window)
	view := &root{
		Grid: tview.NewGrid().
			SetRows(0).
			SetColumns(30, 0),
	}
	view.tree = tree
	view.addTree()
	return view
}

func (view *root) addTree() {
	view.
		AddItem(view.tree.TreeView, 0, 0, 1, 1, 0, 0, true)
}

func (view *root) removeTree() {
	view.RemoveItem(view.tree.TreeView)
}

func (view *root) ShowFileInfo(path string) {
	if view.fileInfo.view != nil {
		view.RemoveItem(view.fileInfo.view)
	}

	fileInfo := newFileInfo(path)
	view.AddItem(fileInfo.view, 0, 1, 1, 1, 0, 0, false)
	view.fileInfo = fileInfo
}

func (view *root) ShowFeedback(text string) {
	view.RemoveFeedback()

	feedback := newFeedback(text)
	view.AddItem(feedback.view, 1, 0, 1, 2, 0, 0, false)
	view.feedback = feedback

	feedback.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		view.RemoveItem(view.feedback.view)
		return nil
	})
}

func (view *root) RemoveFeedback() {
	if view.feedback.view != nil {
		view.RemoveItem(view.feedback.view)
		view.feedback.view = nil
	}
}

func (view *root) showSeach() {
	view.removeSearch()

	search := newSearch(view.tree, view.window)
	view.AddItem(search.view, 1, 0, 1, 1, 0, 0, true)
	SharedConfig.Application.SetFocus(search.view)
	view.search = search
}

func (view *root) removeSearch() {
	if view.search.view != nil {
		view.RemoveItem(view.search.view)
		view.search.view = nil
	}
	SharedConfig.Application.SetFocus(view.tree.TreeView)
}

func (view *root) displayedFeedback() bool {
	return view.feedback.view != nil
}

func (view *root) displayedSearch() bool {
	return view.search.view != nil
}
