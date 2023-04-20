package directory

import "os"

// Exists checks if the directory with the specified name exists.
func Exists(name string) bool {
	fi, err := os.Stat(name)

	if os.IsNotExist(err) {
		return false
	}

	return fi.IsDir()
}
