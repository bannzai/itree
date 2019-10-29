package main

import (
	"github.com/bannzai/itree/pkg/parser"
	"github.com/bannzai/itree/pkg/ui"
	"github.com/rivo/tview"
)

func main() {
	size, err := parser.ParseSize()
	if err != nil {
		panic(err)
	}

	application := tview.NewApplication()
	ui.SharedConfig.Application = application
	if err := application.SetRoot(ui.NewWindow(size.Width, size.Height), true).Run(); err != nil {
		panic(err)
	}
}
