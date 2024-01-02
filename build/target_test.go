// Copyright 2023-2024 Tobias Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestOutPath(t *testing.T) {
	target := Target{
		Name: "myname",
		Os:   "windows",
		Arch: "amd64",
	}

	outPath := target.OutPath()
	outPathParts := strings.Split(outPath, string(os.PathSeparator))
	outPathPartsLen := len(outPathParts)

	if outPathPartsLen != 2 {
		t.Fatalf("Expected %d element(s) but got %d", 2, outPathPartsLen)
	}

	if outPathParts[0] != "dist" {
		t.Errorf("Expected directory `dist` but got `%s`", outPathParts[0])
	}

	if outPathParts[1] != fmt.Sprintf("%s-%s", target.Os, target.Arch) {
		t.Errorf("Expected directory `%s-%s` but got `%s`", target.Os, target.Arch, outPathParts[1])
	}
}

func TestOutFileName(t *testing.T) {
	target := Target{
		Name: "myname",
		Os:   "windows",
		Arch: "amd64",
	}

	outPath := target.OutFileName()
	outPathParts := strings.Split(outPath, string(os.PathSeparator))
	outPathPartsLen := len(outPathParts)

	if outPathPartsLen != 3 {
		t.Fatalf("Expected %d element(s) but got %d", 3, outPathPartsLen)
	}

	if outPathParts[0] != "dist" {
		t.Errorf("Expected directory `dist` but got `%s`", outPathParts[0])
	}

	if outPathParts[1] != fmt.Sprintf("%s-%s", target.Os, target.Arch) {
		t.Errorf("Expected directory `%s-%s` but got `%s`", target.Os, target.Arch, outPathParts[1])
	}

	if outPathParts[2] != target.Name {
		t.Errorf("Expected filename `%s` but got `%s`", target.Name, outPathParts[2])
	}
}
