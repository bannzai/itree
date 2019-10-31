package file

import "os"

func MakeDirectory(path string) error {
	return os.Mkdir(path, 0777)
}
