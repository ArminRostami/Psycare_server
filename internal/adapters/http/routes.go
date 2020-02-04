package http

import (
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

// Handler is an http handler impl using go-chi
type Handler struct {
	Auth   *jwtauth.JWTAuth
	Router *chi.Mux
	*app.UserService
	Validate *validator.Validate
}

// SetupRoutes _
func (h *Handler) SetupRoutes() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Use(render.SetContentType(render.ContentTypeJSON))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", h.createUser)
	usersRouter.Post("/auth", h.login)

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", usersRouter)
	})
}

// Serve _
func (h *Handler) Serve() {
	h.SetupRoutes()
	log.Print("listening on port 5555...")
	log.Fatal(http.ListenAndServe(":5555", h.Router))
}

func renderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, render.M{"data": data, "status": http.StatusOK})
}

func renderError(w http.ResponseWriter, r *http.Request, e *httpError) {
	render.Status(r, e.status)
	render.JSON(w, r, render.M{"type": e.errType, "message": e.err.Error(), "status": e.status})
}
