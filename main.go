package main

import (
	"github.com/bannzai/itree/pkg/ui"
	"github.com/rivo/tview"
)

func main() {
	if err := tview.NewApplication().SetRoot(ui.TreeUI(), true).Run(); err != nil {
		panic(err)
	}
}
