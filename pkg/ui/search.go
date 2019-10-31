package ui

import (
	"github.com/gdamore/tcell"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type search struct {
	view *tview.InputField
}

func newSearch(tree tree, window *Window) search {
	search := search{
		view: tview.
			NewInputField(),
	}

	originalRootPointer := tree.GetRoot()
	copiedRoot := *originalRootPointer
	copiedRootPointer := &copiedRoot
	tree.addNodeAll(copiedRootPointer)
	tree.expandAll(copiedRootPointer)

	search.
		view.
		SetLabel("/").
		SetChangedFunc(func(text string) {
			if len(text) > 0 {
				tree.setAllDisplayTextToPath(copiedRootPointer)
				tree.SetRoot(copiedRootPointer)
				nodes := searchTree(tree, text)
				copiedRootPointer.SetChildren(nodes)
			} else {
				tree.setAllDisplayTextToBasename(originalRootPointer)
				tree.SetRoot(originalRootPointer)
				window.removeSearch()
			}
		}).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				window.removeSearch()
			}
		})
	return search
}

func searchTree(tree tree, text string) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}
	for _, node := range lastNodes(tree.originalRootNode) {
		path := extractNodeReference(node).path
		if fuzzy.Match(text, path) {
			nodes = append(nodes, node)
		}
	}

	return nodes
}
