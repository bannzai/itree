package ui

import "github.com/rivo/tview"

type FileInfo struct {
	*tview.List
}

const (
	// FIXME: How to set empty rune
	noShortcut      = '≠'
	noSecondaryText = ""
)

func NewFileInfo(path string) FileInfo {
	fileInfo := FileInfo{
		tview.NewList().
			AddItem("", noSecondaryText, noShortcut, nil),
	}

	return fileInfo
}
