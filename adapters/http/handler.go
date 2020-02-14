package http

import (
	"log"
	"net/http"
	app "psycare/application"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-playground/validator"
	"github.com/pkg/errors"
)

type Handler struct {
	*app.Services
	Auth     *jwtauth.JWTAuth
	Router   *chi.Mux
	Validate *validator.Validate
}

func (h *Handler) Serve(port string) error {
	h.Router = chi.NewRouter()

	h.SetupRoutes()

	h.Router.Post("/upload/{adv_id}", uploadFile)

	err := h.setupFileServer("/", "static")
	if err != nil {
		return errors.Wrap(err, "failed to serve http")
	}

	err = h.setupFileServer("/files", "files")
	if err != nil {
		return errors.Wrap(err, "failed to serve http")
	}

	log.Printf("listening on port %s...", port)
	return http.ListenAndServe(":"+port, h.Router)
}
