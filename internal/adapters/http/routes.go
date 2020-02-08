package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"

	app "psycare/internal/application"
)

type httpError struct {
	errType string
	status  int
	err     error
}

type Handler struct {
	*app.Services
	Auth     *jwtauth.JWTAuth
	Router   *chi.Mux
	Validate *validator.Validate
}

func (h *Handler) SetupRoutes() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Use(render.SetContentType(render.ContentTypeJSON))

	h.Router.Route("/api/v1", func(r chi.Router) {
		// public routes
		r.Group(func(r chi.Router) {
			r.Post("/users", h.createUser)
			r.Post("/users/auth", h.login)
			r.Get("/advisors", h.getAdvisors)

		})
		// authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(h.Auth))
			r.Use(jwtauth.Authenticator)
			r.Post("/advisors", h.createAdvisor)
			r.Post("/appointments", h.makeAppointment)
		})

	})
}

func (h *Handler) Serve() {
	h.SetupRoutes()
	log.Print("listening on port 5555...")
	log.Fatal(http.ListenAndServe(":5555", h.Router))
}

func (h *Handler) decodeAndValidate(r *http.Request, dst interface{}) *httpError {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request decoding error", err: err}
	}
	err = h.Validate.Struct(dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request validation error", err: err}
	}
	return nil
}

func renderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, render.M{"data": data, "status": http.StatusOK})
}

func renderError(w http.ResponseWriter, r *http.Request, e *httpError) {
	render.Status(r, e.status)
	render.JSON(w, r, render.M{"type": e.errType, "message": e.err.Error(), "status": e.status})
}
