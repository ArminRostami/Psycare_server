package http

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

func FileServer(r chi.Router, path string, root http.FileSystem) error {
	if strings.ContainsAny(path, "{}*") {
		return errors.New("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

	return nil
}

func (h *Handler) setupFileServer(path, dirName string) error {
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "failed to setup routes")
	}
	root := filepath.Join(wd, dirName)
	log.Printf("file server starting for root: %s", root)

	err = FileServer(h.Router, path, http.Dir(root))
	if err != nil {
		return errors.Wrap(err, "failed to setup routes")
	}

	return nil
}
