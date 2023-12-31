// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package directory

import (
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	testdata := map[string]bool{
		"./testdata/testfolder":     true,
		"./testdata/not-a-folder":   false,
		"./testdata/i-do-not-exist": false,
	}

	for dir, exists := range testdata {
		absdir, err := filepath.Abs(dir)
		if err != nil {
			t.Errorf("Failed to obtain absolute path of %s", dir)
		}

		if Exists(absdir) != exists {
			t.Errorf("Exists %s: expected %t but got %t", absdir, exists, !exists)
		}
	}
}
