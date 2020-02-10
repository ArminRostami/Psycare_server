package http

import (
	"log"
	"net/http"
	app "psycare/application"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/go-playground/validator"
)

type Handler struct {
	*app.Services
	Auth     *jwtauth.JWTAuth
	Router   *chi.Mux
	Validate *validator.Validate
}

func (h *Handler) Serve(port string) error {
	h.SetupRoutes()
	log.Printf("listening on port %s...", port)
	return http.ListenAndServe(":"+port, h.Router)
}
