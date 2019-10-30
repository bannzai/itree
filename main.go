package main

import (
	"flag"
	"fmt"

	"github.com/bannzai/itree/pkg/parser"
	"github.com/bannzai/itree/pkg/ui"
	"github.com/rivo/tview"
)

var (
	path = flag.String("path", "./", "specified start path. default is ./")
	help = flag.Bool("help", false, "displayed help message for itree")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Printf(`itree displayed file system tree and command interactively about file system.

Usage:
	itree [options]

The options are:
	--path=$PATH specified start path. default is ./
	--help       displayed help message for itree
`,
		)
		return
	}

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
