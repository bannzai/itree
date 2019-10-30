package ui

import (
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type tree struct {
	*tview.TreeView
	window *Window
}

func newTree(window *Window) tree {
	rootDir := SharedConfig.RootPath
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed).
		SetReference(newNodeReference(rootDir, true, nil))

	tree := tree{
		TreeView: tview.NewTreeView().
			SetRoot(root).
			SetCurrentNode(root),
		window: window,
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
		if tree.window.displayedFeedback() {
			tree.window.RemoveFeedback()
			return nil
		}
		tree.handleEventWithKey(event)
		return event
	})

	return tree
}

func (tree tree) name() string {
	return nameOfTree
}

func (tree tree) view() tview.Primitive {
	return tree.TreeView
}

func (tree *tree) addNode(directoryNode *tview.TreeNode, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		node := createTreeNode(file.Name(), file.IsDir(), directoryNode)
		directoryNode.AddChild(node)
	}
}

func (tree *tree) handleEventWithKey(event *tcell.EventKey) {
	switch event.Rune() {
	case 'c':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		path := nodeReference.path
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
		tree.window.ShowFeedback(fmt.Sprintf("copy relative path to clipboard %s", path))
	case 'C':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		path := absolutePath(*nodeReference)
		if err := clipboard.WriteAll(path); err != nil {
			fmt.Printf("clipboard.WriteAll(%s) is error. error is %v", path, err)
			return
		}
		tree.window.ShowFeedback(fmt.Sprintf("copy absolute path to clipboard %s", path))
	case 'r':
		tree.window.SwitchRenameForm(tree.GetCurrentNode())
	case 'o':
		path := extractNodeReference(tree.GetCurrentNode()).path
		if err := exec.Command("open", path).Run(); err != nil {
			panic(fmt.Errorf("open %s is error, raw error %w", path, err))
		}
		tree.window.ShowFeedback(fmt.Sprintf("$ open %s", path))
	case 'n':
		tree.window.SwitchAddFileForm(tree.GetCurrentNode())
	case 'N':
		tree.window.SwitchAddDirectoryForm(tree.GetCurrentNode())
	case 'e':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		NewEditor().Launch(nodeReference.path)
	case 'i':
		nodeReference := extractNodeReference(tree.GetCurrentNode())
		tree.window.ShowFileInfo(nodeReference.path)
	case '?':
		tree.window.SwitchUsage()
	}
}
