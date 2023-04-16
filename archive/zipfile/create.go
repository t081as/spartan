package zipfile

import (
	"archive/zip"
	"io"
	"os"
)

// Create creates a zip file containing the given files.
//
//	files := map[string]string{
//	  "./testdata/create-1.txt": "f1.txt",
//	  "./testdata/create-2.txt": "data/f2.txt",
//	}
//
//	err := Create("test.zip", files)
func Create(filename string, files map[string]string) error {
	zipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for sourcefilename, destinationfilename := range files {
		if err := addFileToZipWriter(zipWriter, sourcefilename, destinationfilename); err != nil {
			return err
		}
	}

	return nil
}

func addFileToZipWriter(zipWriter *zip.Writer, sourcefilename, destinationfilename string) error {
	sourcefile, err := os.Open(sourcefilename)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	info, err := sourcefile.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = destinationfilename
	header.Method = zip.Deflate

	fileWriter, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, sourcefile)
	return err
}
