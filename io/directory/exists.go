// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
