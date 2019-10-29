package ui

import "github.com/rivo/tview"

type Config struct {
	*tview.Application
}

var SharedConfig = Config{
	Application: nil,
}
