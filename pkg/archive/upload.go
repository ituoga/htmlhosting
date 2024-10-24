package archive

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadFile uploads a ZIP file to the server
func UploadFile(url, filePath string) error {
	// Open the ZIP file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffer and multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create the file part in the multipart form
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return err
	}

	// Copy the file content to the multipart form part
	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the writer to finalize the multipart form
	err = writer.Close()
	if err != nil {
		return err
	}

	// Create a new HTTP request to upload the file
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform the HTTP request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Read and print the response
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))

	return nil
}
