package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rivo/tview"
)

type fileInfo struct {
	view *tview.TextView
}

func newFileInfo(path string) fileInfo {
	textView := tview.NewTextView().
		SetText(buildFileInfomation(path))
	textView.
		SetBorder(true).
		SetTitle("File Info")

	fileInfo := fileInfo{
		view: textView,
	}

	return fileInfo
}

func buildFileInfomation(path string) string {
	fileStat, err := os.Stat(path)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf(`
	File Name     : %s
	Size          : %d bytes
	Permissions   : %d
	Last Modified : %s
	Is Directory  : %t
	`,
		fileStat.Name(),
		fileStat.Size(),
		fileStat.Mode(),
		formatLastModified(fileStat.ModTime()),
		fileStat.IsDir(),
	)
}

func formatLastModified(time time.Time) string {
	layout := "2006-01-02 15::03::04 (Mon)"
	return time.Format(layout)
}
