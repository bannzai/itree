package ui

import (
	"github.com/gdamore/tcell"
)

func (window *Window) SwitchSearchView() {
	tree := NewTree(window)
	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		tree.handlerInputEventWithKey(event)
		return nil
	})

	window.Root.AddAndSwitchToPage(nameOfSearch, tree, true)
}

func (tree *Tree) handlerInputEventWithKey(event *tcell.EventKey) {

}
