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
	SwitchUsage()
	ShowFileInfo(path string)
	ShowFeedback(text string)
	RemoveFeedback()
	displayedFeedback() bool
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
		if tree.switcher.displayedFeedback() {
			tree.switcher.RemoveFeedback()
			return nil
		}
		tree.handleEventWithKey(event)
		return event
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
		path := nodeReference.path
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
	case 'C':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		path := absolutePath(*nodeReference)
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
		tree.switcher.ShowFeedback(fmt.Sprintf("copy to clipboard %s", path))
	case 'r':
		tree.switcher.SwitchRenameForm(tree.GetCurrentNode())
	case 'o':
		path := extractNodeReference(tree.GetCurrentNode()).path
		if err := exec.Command("open", path).Run(); err != nil {
			panic(fmt.Errorf("open %s is error, raw error %w", path, err))
		}
	case 'n':
		tree.switcher.SwitchAddFileForm(tree.GetCurrentNode())
	case 'N':
		tree.switcher.SwitchAddDirectoryForm(tree.GetCurrentNode())
	case 'e':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		NewEditor().Launch(nodeReference.path)
	case 'i':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		tree.switcher.ShowFileInfo(nodeReference.path)
	case '?':
		tree.switcher.SwitchUsage()
	}
}
