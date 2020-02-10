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

func (h *Handler) Serve() {
	h.SetupRoutes()
	log.Print("listening on port 5555...")
	log.Fatal(http.ListenAndServe(":5555", h.Router))
}
