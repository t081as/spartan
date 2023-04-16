package file

import (
	"errors"
	"os"
	"path/filepath"
)

// RemoveGlob removes files using the given glob pattern.
func RemoveGlob(pattern string) error {
	files, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	var errs []error

	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
