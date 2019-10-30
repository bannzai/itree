package ui

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type switcher interface {
	SwitchRenameForm(node *tview.TreeNode)
	SwitchAddFileForm(*tview.TreeNode)
	SwitchAddDirectoryForm(*tview.TreeNode)
}

type Tree struct {
	*tview.TreeView
	switcher
}

func NewTree(switcher switcher) Tree {
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed).
		SetReference(newNodeReference(rootDir, true, nil))

	tree := Tree{
		TreeView: tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root),
		switcher: switcher,
	}

	tree.addNode(root, rootDir)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return
		}
		nodeReference := reference.(*nodeReference)
		if !nodeReference.isDir {
			return
		}

		children := node.GetChildren()
		if len(children) == 0 {
			path := nodeReference.path
			tree.addNode(node, path)
		} else {
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

func (tree *Tree) addNode(directoryNode *tview.TreeNode, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := createTreeNode(file.Name(), file.IsDir(), directoryNode)
		directoryNode.AddChild(node)
	}
}

func (tree *Tree) handleEventWithKey(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		path := absolutePath(*nodeReference)
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
	case 'r':
		tree.switcher.SwitchRenameForm(tree.GetCurrentNode())
	case 'o':
		path := extractNodeReference(tree.GetCurrentNode()).path
		if err := exec.Command("open", path).Run(); err != nil {
			panic(err)
		}
	case 'n':
		tree.switcher.SwitchAddFileForm(tree.GetCurrentNode())
	case 'N':
		tree.switcher.SwitchAddDirectoryForm(tree.GetCurrentNode())
	case 'e':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		NewEditor().Launch(nodeReference.path)
	}
}
