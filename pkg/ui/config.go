package ui

import "github.com/rivo/tview"

type Config struct {
	RootPath string
	*tview.Application
}

var SharedConfig = Config{
	Application: nil,
}
