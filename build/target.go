package build

import (
	"fmt"
	"path/filepath"
)

// Target defines a possible compilation target.
type Target struct {
	Name        string // Application name
	Os          string // Operating system (GOOS)
	Architcture string // Architecture (GOARCH)
}

// OutPath returns a relative path to a operating system and architecture
// specific directory in a `dist` subdirectory.
func (t *Target) OutPath() string {
	return filepath.Join(".", "dist", fmt.Sprintf("%s-%s", t.Os, t.Architcture))
}

// OutFileName combines OutPath with the application name
func (t *Target) OutFileName() string {
	return filepath.Join(t.OutPath(), t.Name)
}
