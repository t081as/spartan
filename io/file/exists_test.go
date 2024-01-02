// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	testdata := map[string]bool{
		"./testdata/.gitkeep":           true,
		"./testdata/i-do-not-exist.txt": false,
	}

	for file, exists := range testdata {
		filename, err := filepath.Abs(file)
		if err != nil {
			t.Errorf("Failed to obtain absolute path of %s", file)
		}

		if Exists(filename) != exists {
			t.Errorf("Exists %s: expected %t but got %t", filename, exists, !exists)
		}
	}
}
