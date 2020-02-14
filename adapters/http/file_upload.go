package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "adv_id")

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		err := errors.Wrap(err, "Error setting parse type")
		renderError(w, r, &httpError{"file upload failed", http.StatusInternalServerError, err})
		return
	}

	file, handler, err := r.FormFile("fileUpload")
	if err != nil {
		err := errors.Wrap(err, "Error Retrieving the File")
		renderError(w, r, &httpError{"file upload failed", http.StatusInternalServerError, err})
		return
	}
	defer file.Close()
	log.Printf("Uploaded File: %+v\n", handler.Filename)
	log.Printf("File Size: %+v\n", handler.Size)

	wd, err := os.Getwd()
	if err != nil {
		err := errors.Wrap(err, "failed to get working directory")
		renderError(w, r, &httpError{"file upload failed", http.StatusInternalServerError, err})
		return
	}

	dir := filepath.Join(wd, "files", fmt.Sprintf("adv_%s.png", id))
	f, err := os.Create(dir)
	if err != nil {
		err := errors.Wrap(err, "failed to create file")
		renderError(w, r, &httpError{"file upload failed", http.StatusInternalServerError, err})
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		err := errors.Wrap(err, "failed to copy file to server")
		renderError(w, r, &httpError{"file upload failed", http.StatusInternalServerError, err})
		return
	}

	renderData(w, r, "Successfully Uploaded File")
}
