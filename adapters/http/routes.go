package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func (h *Handler) SetupRoutes() {

	h.Router.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// public routes
			r.Post("/users", h.createUser)
			r.Post("/users/auth", h.login)
			r.Get("/advisors", h.getAdvisors)
			r.Get("/advisors/{adv_id}", h.getAdvisor)
			r.Get("/advisors/{adv_id}/rating", h.getAvgRating)
			r.Get("/advisors/schedule/{adv_id}", h.getScheduleWithID)
			r.Get("/appointments/advisor/{adv_id}", h.getAppointmentsWithID)

		})
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(h.Auth))
			r.Use(jwtauth.Authenticator)
			// authenticated routes
			r.Get("/users", h.getUser)
			r.Post("/advisors", h.createAdvisor)
			r.Post("/advisors/schedule", h.addSchedule)
			r.Get("/advisors/schedule", h.getSchedule)
			r.Delete("/advisors/schedule", h.deleteSchedule)
			r.Post("/appointments", h.bookAppointment)
			r.Get("/appointments/user", h.getUserAppointments)
			r.Get("/appointments/advisor", h.getAdvisorAppointments)
			r.Post("/appointments/rate", h.rateAppointment)
			r.Post("/appointments/cancel", h.cancelAppointment)
		})
	})
}
