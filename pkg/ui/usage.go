package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Usage struct {
	*tview.List
}

func NewUsage(window *Window) Usage {
	usage := Usage{}
	list := tview.NewList().
		ShowSecondaryText(false).
		SetSelectedTextColor(tview.Styles.PrimaryTextColor).
		SetSelectedBackgroundColor(tview.Styles.PrimitiveBackgroundColor).
		AddItem("Copy selected node file path", "", 'c', nil).
		AddItem("Copy selected node absolute file path", "", 'C', nil).
		AddItem("Rename file node", "", 'r', nil).
		AddItem("'$ open $FILE_PATH'", "", 'o', nil).
		AddItem("New file", "", 'n', nil).
		AddItem("New directory", "", 'N', nil).
		AddItem("Open current node with $EDITOR. Default is vim", "", 'e', nil).
		AddItem("appear information for current node", "", 'i', nil).
		AddItem("change mode for file search", "", '/', nil).
		AddItem("help message for usage itree", "", '?', nil)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		window.transition.RemovePage(usage.name())
		return nil
	})

	usage.List = list
	return usage
}

func (window *Window) SwitchUsage() {
	usage := NewUsage(window)
	window.AddAndSwitchToPage(usage.name(), usage.view(), true)
}

func (Usage) name() string {
	return nameOfUsage
}

func (usage Usage) view() tview.Primitive {
	return usage.List
}
