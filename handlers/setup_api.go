package handlers

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/ituoga/htmlhosting/pkg/archive"
)

func SetupApi(logger *slog.Logger, router chi.Router, session sessions.Store, ns *embeddednats.Server) error {

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
			// Allow only POST requests
			if r.Method != http.MethodPost {
				http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
				return
			}

			// Parse the multipart form to extract the file
			err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
			if err != nil {
				http.Error(w, "Error parsing form data", http.StatusBadRequest)
				return
			}

			// Retrieve the file from form data
			file, handler, err := r.FormFile("file")
			if err != nil {
				http.Error(w, "Failed to read file from form data", http.StatusBadRequest)
				return
			}
			defer file.Close()

			// Save the file to a temporary location
			tempDir := "./uploads"
			os.MkdirAll(tempDir, os.ModePerm)
			tempFilePath := filepath.Join(tempDir, handler.Filename)

			tempFile, err := os.Create(tempFilePath)
			if err != nil {
				http.Error(w, "Unable to create temporary file", http.StatusInternalServerError)
				return
			}
			defer tempFile.Close()

			// Copy the uploaded file to the temporary location
			_, err = io.Copy(tempFile, file)
			if err != nil {
				http.Error(w, "Failed to save file", http.StatusInternalServerError)
				return
			}

			// Decompress the file if it's a ZIP file
			if filepath.Ext(handler.Filename) == ".zip" {
				fmt.Fprintf(w, "Decompressing ZIP file: %s\n", handler.Filename)
				err := archive.Unzip(tempFilePath, filepath.Join(tempDir, "extracted"))
				if err != nil {
					http.Error(w, "Failed to decompress file", http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "Successfully decompressed ZIP file: %s\n", handler.Filename)
			} else {
				http.Error(w, "Only ZIP files are supported", http.StatusBadRequest)
				return
			}

			fmt.Fprintf(w, "File upload and decompression successful")
		})
	})

	return nil
}
