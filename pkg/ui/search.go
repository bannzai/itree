package ui

import (
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/rivo/tview"
)

type search struct {
	view *tview.InputField
}

func newSearch(tree tree) search {
	search := search{
		view: tview.
			NewInputField(),
	}

	copiedRoot := *tree.GetRoot()
	copiedRootPointer := &copiedRoot
	tree.SetRoot(copiedRootPointer)

	search.view.SetChangedFunc(func(text string) {
		nodes := searchTree(tree, text)
		copiedRootPointer.SetChildren(nodes)
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
