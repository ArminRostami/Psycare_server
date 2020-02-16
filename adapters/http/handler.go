package http

import (
	"log"
	"net/http"
	app "psycare/application"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
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
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	h.Router.Use(cors.Handler)

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
