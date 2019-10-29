package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type ErrorField struct {
	*tview.TextView
}

func NewErrorField() ErrorField {
	textView := tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	return ErrorField{textView}
}
