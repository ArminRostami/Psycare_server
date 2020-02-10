package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

func (h *Handler) SetupRoutes() {
	h.Router = chi.NewRouter()
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Use(render.SetContentType(render.ContentTypeJSON))

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// public routes
			r.Post("/users", h.createUser)
			r.Post("/users/auth", h.login)
			r.Get("/advisors", h.getAdvisors)

		})
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(h.Auth))
			r.Use(jwtauth.Authenticator)
			// authenticated routes
			r.Post("/advisors", h.createAdvisor)
			r.Post("/advisors/schedule", h.addSchedule)
			r.Post("/appointments", h.makeAppointment)
			r.Get("/appointments/user", h.getUserAppointments)
			r.Get("/appointments/advisor", h.getAdvisorAppointments)
		})

	})
}
