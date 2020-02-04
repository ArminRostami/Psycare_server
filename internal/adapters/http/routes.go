package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"

	app "psycare/internal/application"
)

// Handler is an http handler impl using go-chi
type Handler struct {
	Router   *chi.Mux
	Us       *app.UserService
	Validate *validator.Validate
}

func (h *Handler) SetupRoutes() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Use(render.SetContentType(render.ContentTypeJSON))

	usersRouter := chi.NewRouter()
	usersRouter.Post("/", h.addUser)

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/users", usersRouter)
	})
}

func (h *Handler) Serve() {
	h.SetupRoutes()
	log.Print("listening on port 5555...")
	log.Fatal(http.ListenAndServe(":5555", h.Router))
}

func renderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, render.M{"data": data, "status": http.StatusOK})
}

func renderError(w http.ResponseWriter, r *http.Request, status int, errType string, err error) {
	render.Status(r, status)
	render.JSON(w, r, render.M{"type": errType, "message": err.Error(), "status": status})
}
