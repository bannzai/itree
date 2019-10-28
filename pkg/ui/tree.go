package ui

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type Tree struct {
	*tview.TreeView
	switcher
}

type nodeReference struct {
	path       string
	isDir      bool
	parentNode *tview.TreeNode
}

func (reference *nodeReference) setPath(path string) {
	reference.path = path
}

func newNodeReference(path string, isDir bool, parentNode *tview.TreeNode) *nodeReference {
	return &nodeReference{
		path:       path,
		isDir:      isDir,
		parentNode: parentNode,
	}
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

	tree.addNode(root, rootDir)

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return
		}
		nodeReference := reference.(nodeReference)
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
		if file.IsDir() {
			node.SetColor(tcell.ColorGreen)
		}
		directoryNode.AddChild(node)
	}
}

func (tree *Tree) handleEventWithKey(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		path := absolutePath(*tree.GetCurrentNode().GetReference().(*nodeReference))
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
	case 'r':
		tree.switcher.SwitchRenameForm(tree.GetCurrentNode())
	case 'o':
		path := tree.GetCurrentNode().GetReference().(nodeReference).path
		if err := exec.Command("open", path).Run(); err != nil {
			panic(err)
		}
	case 'n':
		tree.switcher.SwitchAddFileForm(tree.GetCurrentNode())
	}
}
