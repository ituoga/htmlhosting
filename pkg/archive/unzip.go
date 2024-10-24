package archive

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Unzip decompresses the specified ZIP file to the target directory
func Unzip(zipFilePath, targetDir string) error {
	// Open the ZIP file
	zipReader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Create target directory if it doesn't exist
	err = os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		return err
	}

	// Iterate through each file in the ZIP archive
	for _, file := range zipReader.File {
		// Construct the full path for the file or directory in the target directory

		filePath := filepath.Join(targetDir, file.Name)
		log.Printf("Extracting file: %s\n", filePath)
		// Check if the file is a directory
		if file.FileInfo().IsDir() {
			// Create the directory
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// Create the file's directory if it doesn't exist
		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		// Create the file
		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer destFile.Close()

		// Open the zipped file for reading
		zippedFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zippedFile.Close()

		// Copy the content of the zipped file to the new file
		_, err = io.Copy(destFile, zippedFile)
		if err != nil {
			return err
		}
	}

	return nil
}
