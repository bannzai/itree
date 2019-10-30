package ui

import "github.com/rivo/tview"

type feedback struct {
	view *tview.TextView
}

func newFeedback(text string) feedback {
	textView := tview.NewTextView().
		SetText(text)
	textView.
		SetBorder(true)

	feedback := feedback{
		view: textView,
	}

	return feedback
}
