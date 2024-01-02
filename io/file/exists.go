// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
