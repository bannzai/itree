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
}

func NewTree() Tree {
	rootDir := "./"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := Tree{
		tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root),
	}

	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
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
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			add(node, path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	defaultInputHandler := tree.InputHandler()
	defaultInputCapture := tree.GetInputCapture()
	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !containsCaptureKey(event) {
			return defaultInputCapture(event)
		}

		tree.handleEventWithKey(event)
		return event
	})

	tree.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		defaultInputHandler(event, setFocus)
	})

	return tree
}

func (tree Tree) name() string {
	return nameOfTree
}

func (tree Tree) view() tview.Primitive {
	return tree.TreeView
}

func containsCaptureKey(event *tcell.EventKey) bool {
	if event.Key() != tcell.KeyRune {
		return false
	}

	switch event.Rune() {
	case 'c', 'r':
		return true
	default:
		return false
	}
}

func (tree *Tree) handleEventWithKey(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		path := tree.GetCurrentNode().GetReference().(string)
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
	case 'r':
		path := tree.GetCurrentNode().GetReference().(string)
		fmt.Println(path)
	}
}
