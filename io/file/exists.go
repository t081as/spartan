package file

import "os"

// Exists checks if the file with the given filename exists.
func Exists(filename string) bool {
	fi, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !fi.IsDir()
}
