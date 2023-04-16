package tgzfile

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

// Create creates a tar.gz file containing the given files.
// File permissions are mirrored from the file system.
//
//	files := map[string]string{
//	  "./testdata/create-1.txt": "execute-me", // source file name : destination file name
//	  "./testdata/create-2.txt": "data/f2.txt",
//	}
//
//	err := Create("test.tar.gz", files)
func Create(filename string, files map[string]string) error {
	permissions := func(string) int64 {
		return -1
	}

	return CreatePermissionFunc(filename, files, permissions)
}

// CreatePermissionFunc creates a tar.gz file containing the given files.
// File permissions may be changed by the permission func.
//
//	files := map[string]string{
//	  "./testdata/create-1.txt": "execute-me", // source file name : destination file name
//	  "./testdata/create-2.txt": "data/f2.txt",
//	}
//
//	permissions := func(filename string) int64 {
//	  if filename == "execute-me" {
//	    return 0777 // Overwrite permissions
//	  }
//
//	  return -1 // Use permissions from file system
//	}
//
//	err := CreatePermissionFunc("test.tar.gz", files, permissions)
func CreatePermissionFunc(filename string, files map[string]string, permissions func(string) int64) error {
	gzFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer gzFile.Close()

	gzWriter := gzip.NewWriter(gzFile)
	defer gzWriter.Close()

	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	for sourcefilename, destinationfilename := range files {
		filemode := int64(-1)

		if permissions != nil {
			filemode = permissions(destinationfilename)
		}

		if err := addFileToTarWriter(tarWriter, sourcefilename, destinationfilename, filemode); err != nil {
			return err
		}
	}

	return nil
}

func addFileToTarWriter(tarWriter *tar.Writer, sourcefilename, destinationfilename string, mode int64) error {
	sourcefile, err := os.Open(sourcefilename)
	if err != nil {
		return err
	}
	defer sourcefile.Close()

	info, err := sourcefile.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, "")
	if err != nil {
		return err
	}

	header.Name = destinationfilename

	if mode != -1 {
		header.Mode = mode
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, sourcefile)
	return err
}
