package fetcher

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Size struct {
	Width, Height int
}

func ParseSize() (Size, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		return Size{}, errors.Wrap(err, "stty size is failed")
	}
	size := string(output)
	return parseSize(size), nil
}

func parseSize(sizeString string) Size {
	size := Size{}
	converted := func(component string) int {
		edge, err := strconv.Atoi(component)
		if err != nil {
			panic(err)
		}
		return edge
	}
	for i, component := range strings.Split(sizeString, " ") {
		if i == 0 {
			size.Width = converted(component)
			continue
		}
		if i == 1 {
			size.Height = converted(strings.ReplaceAll(component, "\n", ""))
			continue
		}
		panic(fmt.Sprintf("Unexpected index %d, for size string of %s", i, sizeString))
	}
	return size
}
