package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type root struct {
	*tview.Grid
	fileInfo
	feedback

	isDisplayedFeedback bool
}

func newRoot(window *Window) *root {
	tree := newTree(window)
	view := &root{
		Grid: tview.NewGrid().
			SetRows(0, 5).
			SetColumns(30, 0).
			AddItem(tree, 0, 0, 1, 1, 0, 0, true),
	}
	return view
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
	view.AddItem(feedback.view, 1, 0, 1, 2, 0, 0, true)
	view.feedback = feedback

	feedback.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		view.RemoveItem(view.feedback.view)
		return nil
	})

	view.isDisplayedFeedback = true
}

func (view *root) RemoveFeedback() {
	if view.feedback.view != nil {
		view.RemoveItem(view.feedback.view)
	}
	view.isDisplayedFeedback = false
}

func (view *root) displayedFeedback() bool {
	return view.isDisplayedFeedback
}
