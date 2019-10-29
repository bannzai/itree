package ui

import "github.com/rivo/tview"

type ErrorField struct {
	*tview.TextView
}

func NewErrorField() ErrorField {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	return ErrorField{textView}
}
