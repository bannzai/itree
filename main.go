package main

import (
	"flag"

	"github.com/bannzai/itree/pkg/parser"
	"github.com/bannzai/itree/pkg/ui"
	"github.com/rivo/tview"
)

var (
	path = flag.String("path", "./", "Root path. Default is ./")
)

func main() {
	flag.Parse()

	size, err := parser.ParseSize()
	if err != nil {
		panic(err)
	}

	application := tview.NewApplication()
	ui.SharedConfig.Application = application
	ui.SharedConfig.RootPath = *path
	if err := application.SetRoot(ui.NewWindow(size.Width, size.Height), true).Run(); err != nil {
		panic(err)
	}
}
