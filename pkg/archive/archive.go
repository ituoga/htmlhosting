package archive

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipDirectory compresses the specified folder into a ZIP file
func ZipDirectory(sourceDir, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}

		// If it's a directory, create a directory entry in the ZIP file
		if info.IsDir() {
			if relativePath != "." {
				_, err := zipWriter.Create(relativePath + "/")
				return err
			}
			return nil
		}

		// Open the file to read its content
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create the file entry in the ZIP archive
		writer, err := zipWriter.Create(relativePath)
		if err != nil {
			return err
		}

		// Copy the file content to the ZIP entry
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
