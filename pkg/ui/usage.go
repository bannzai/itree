package ui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Usage struct {
	*tview.TextView
}

func NewUsage(window *Window) Usage {
	view := Usage{}
	view.TextView = tview.NewTextView().
		SetText(usage()).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				window.RemovePage(view.name())
			}
		})
	return view
}

func usage() string {
	return `USAGE 
'c' copy selected node file path, 'C' copy selected node absolute file path. 
'r' rename file node. o '$ open $FILE_PATH'
'n' new file, 'N' new directory, under the selected node.
'e' open current node with $EDITOR. default is vim. 
'i' appear information for current node
	`
}

func (window *Window) SwitchUsage() {
	usage := NewUsage(window)
	window.AddAndSwitchToPage(usage.name(), usage.view(), true)
}

func (Usage) name() string {
	return nameOfUsage
}

func (usage Usage) view() tview.Primitive {
	return usage.TextView
}
