// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package file

import (
	"os"
	"testing"
)

func TestRemoveGlob(t *testing.T) {
	// Create three files
	f1, err := os.Create("./testdata/test-1.txt")
	if err != nil {
		t.Error("Unable to create file test-1.txt")
	}
	f1.Close()

	f2, err := os.Create("./testdata/test-2.txt")
	if err != nil {
		t.Error("Unable to create file test-2.txt")
	}
	f2.Close()

	f3, err := os.Create("./testdata/keep-1.txt")
	if err != nil {
		t.Error("Unable to create file keep-1.txt")
	}
	f3.Close()

	// Shall remove two files
	rmerr := RemoveGlob("./testdata/te*.txt")
	if rmerr != nil {
		t.Errorf("Error while removing files: %v", err)
	}

	if Exists("./testdata/test-1.txt") {
		t.Error("File test-1.txt not removed")
	}

	if Exists("./testdata/test-2.txt") {
		t.Error("File test-2.txt not removed")
	}

	if !Exists("./testdata/keep-1.txt") {
		t.Error("File keep-1.txt removed")
	}

	// Delete file that is not affected by RemoveGlob
	os.Remove("./testdata/keep-1.txt")
}

func TestRemoveGlobWrongPattern(t *testing.T) {
	err := RemoveGlob("[^")

	if err == nil {
		t.Error("Expected error but got nil")
	}
}
