package build

import (
	"fmt"
	"path/filepath"
)

// Target defines a possible compilation target.
type Target struct {
	Name        string // Application name
	Os          string // Operating system (GOOS)
	Arch        string // Architecture (GOARCH)
	Architcture string // Deprecated: use Arch instead
}

// OutPath returns a relative path to a operating system and architecture
// specific directory in a `dist` subdirectory.
func (t *Target) OutPath() string {
	return filepath.Join(".", "dist", fmt.Sprintf("%s-%s", t.Os, t.Arch))
}

// OutFileName combines OutPath with the application name.
func (t *Target) OutFileName() string {
	return filepath.Join(t.OutPath(), t.Name)
}
