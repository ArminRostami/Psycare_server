package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"

	server "psycare/internal/application"
)

type handler struct {
	router   *chi.Mux
	us       *server.UserService
	validate *validator.Validate
}

func (h *handler) setupRoutes() {
	h.router = chi.NewRouter()
	h.router.Use(middleware.Logger)
	h.router.Use(middleware.Recoverer)
	h.router.Use(render.SetContentType(render.ContentTypeJSON))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", h.addUser)

	h.router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", usersRouter)
	})
}

func (h *handler) serve() {
	h.setupRoutes()
	log.Fatal(http.ListenAndServe(":5555", h.router))
}

func renderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, render.M{"data": data, "status": http.StatusOK})
}

func renderError(w http.ResponseWriter, r *http.Request, status int, errType string, err error) {
	render.Status(r, status)
	render.JSON(w, r, render.M{"type": errType, "message": err.Error(), "status": status})
}
