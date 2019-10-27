package ui

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Tree struct {
	*tview.TreeView
	switcher
}

type nodeReference struct {
	path  string
	isDir bool
}

func NewTree(switcher switcher) Tree {
	rootDir := "./"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := Tree{
		TreeView: tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root),
		switcher: switcher,
	}

	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(
					nodeReference{
						path:  filepath.Join(path, file.Name()),
						isDir: file.IsDir(),
					},
				).
				SetSelectable(true)
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	// Add the current directory to the root node.
	add(root, rootDir)

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		nodeReference := reference.(nodeReference)
		if !nodeReference.isDir {
			return
		}

		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := nodeReference.path
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		tree.handleEventWithKey(event)
		return event
	})

	tree.InputHandler()
	tree.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	})

	return tree
}

func (tree Tree) name() string {
	return nameOfTree
}

func (tree Tree) view() tview.Primitive {
	return tree.TreeView
}

func (tree *Tree) handleEventWithKey(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		path := tree.GetCurrentNode().GetReference().(nodeReference).path
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
	case 'r':
		tree.switcher.SwitchRenameForm(tree.GetCurrentNode())
	}
}
