package ui

import "github.com/rivo/tview"

type FormLayout struct {
	*tview.Grid
}

func NewFormLayout(form tview.Primitive, errorField ErrorField) FormLayout {
	return FormLayout{
		tview.NewGrid().
			SetRows(0, 1).
			SetColumns(0).
			AddItem(form, 0, 0, 1, 1, 0, 0, true).
			AddItem(errorField, 1, 0, 1, 1, 0, 30, false),
	}
}
