package ui

import "github.com/rivo/tview"

type Feedback struct {
	View *tview.TextView
}

func NewFeedback(text string) Feedback {
	textView := tview.NewTextView().
		SetText(text)
	textView.
		SetBorder(true)

	feedback := Feedback{
		View: textView,
	}

	return feedback
}
