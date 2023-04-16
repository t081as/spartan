package tgzfile

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"testing"
)

func TestCreatePermissionFunc(t *testing.T) {
	files := map[string]string{
		"./testdata/create-1.txt": "execute-me",
		"./testdata/create-2.txt": "data/f2.txt",
	}

	permissions := func(filename string) int64 {
		if filename == "execute-me" {
			return 0777
		}

		return -1
	}

	err := CreatePermissionFunc("test.tar.gz", files, permissions)
	if err != nil {
		t.Errorf("Error while creating zip archive: %v", err)
	}
	defer os.Remove("test.tar.gz")

	file, err := os.Open("test.tar.gz")
	if err != nil {
		t.Error("Unable to open test.tar.gz (read)")
	}
	defer file.Close()

	gzreader, err := gzip.NewReader(file)
	if err != nil {
		t.Error("Unable to open test.tar.gz (gzip reader)")
	}
	defer gzreader.Close()

	tarreader := tar.NewReader(gzreader)
	filecount := 0

	for {
		header, err := tarreader.Next()

		if err == io.EOF {
			break
		}

		found := false

		for _, originalfile := range files {
			if header.Name == originalfile {
				permission := permissions(originalfile)

				if permission != -1 {
					if permission == header.Mode {
						found = true
					}
				} else {
					found = true
				}
			}
		}

		if !found {
			t.Errorf("test.tar.gz: entry with unexpected values found (%s)", header.Name)
		}

		filecount++
	}

	if filecount != 2 {
		t.Errorf("test.tar.gz: expected 2 files but got %d", filecount)
	}
}
