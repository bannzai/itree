package ui

import (
	"github.com/gdamore/tcell"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type search struct {
	view *tview.InputField
}

func newSearch(tree *tree, window *Window) search {
	search := search{
		view: tview.
			NewInputField(),
	}

	originalRootPointer := tree.GetRoot()
	copiedRoot := *originalRootPointer
	copiedRootPointer := &copiedRoot
	tree.addNodeAll(copiedRootPointer)
	tree.expandAll(copiedRootPointer)
	tree.setAllDisplayTextToPath(copiedRootPointer)
	tree.SetRoot(copiedRootPointer)

	onSearch := func(text string) {
		tree.searchedText = text
		nodes := searchTree(tree, text)
		copiedRootPointer.SetChildren(nodes)
	}

	if len(tree.searchedText) > 0 {
		onSearch(tree.searchedText)
	}

	search.
		view.
		SetLabel("/").
		SetText(tree.searchedText).
		SetChangedFunc(onSearch).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				if len(tree.searchedText) <= 0 {
					tree.setAllDisplayTextToBasename(tree.originalRootNode)
					tree.SetRoot(tree.originalRootNode)
				}
				window.removeSearch()
			}
		})
	return search
}

func searchTree(tree *tree, text string) []*tview.TreeNode {
	nodes := []*tview.TreeNode{}
	for _, node := range lastNodes(tree.originalRootNode) {
		path := extractNodeReference(node).path
		if fuzzy.Match(text, path) {
			nodes = append(nodes, node)
		}
	}

	return nodes
}
