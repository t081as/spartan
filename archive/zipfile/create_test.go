package zipfile

import (
	"archive/zip"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	files := map[string]string{
		"./testdata/create-1.txt": "f1.txt",
		"./testdata/create-2.txt": "data/f2.txt",
	}

	err := Create("test.zip", files)
	if err != nil {
		t.Errorf("Error while creating zip archive: %v", err)
	}
	defer os.Remove("test.zip")

	reader, err := zip.OpenReader("test.zip")
	if err != nil {
		t.Error("Unable to open test.zip (read)")
	}
	defer reader.Close()

	readfiles := reader.File

	if len(readfiles) != 2 {
		t.Errorf("test.zip: expected 2 files but got %d", len(readfiles))
	}

	for _, readfile := range readfiles {
		found := false

		for _, originalfile := range files {
			if readfile.Name == originalfile {
				found = true
			}
		}

		if !found {
			t.Errorf("test.zip: unexpected file %s found", readfile.Name)
		}
	}
}
