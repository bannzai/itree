package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type MainView struct {
	*tview.Grid
	FileInfo
	Feedback

	isDisplayedFeedback bool
}

func NewMainView(window *Window) *MainView {
	tree := NewTree(window)
	view := &MainView{
		Grid: tview.NewGrid().
			SetRows(0, 5).
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

func (view *MainView) ShowFeedback(text string) {
	view.RemoveFeedback()

	feedback := NewFeedback(text)
	view.AddItem(feedback.View, 1, 0, 1, 2, 0, 0, true)
	view.Feedback = feedback

	feedback.View.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		view.RemoveItem(view.Feedback.View)
		return nil
	})

	view.isDisplayedFeedback = true
}

func (view *MainView) RemoveFeedback() {
	if view.Feedback.View != nil {
		view.RemoveItem(view.Feedback.View)
	}
	view.isDisplayedFeedback = false
}

func (view *MainView) displayedFeedback() bool {
	return view.isDisplayedFeedback
}
