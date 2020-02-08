package http

import (
	"net/http"
	"psycare/internal/domain"
)

func (h *Handler) makeAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appt := &domain.Appointment{}
	httpErr = h.decodeAndValidate(r, appt)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	appt.UserID = id
	err := h.CreateAppointment(appt)
	if err != nil {
		renderError(w, r, &httpError{"failed to add appointment", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, "appointment added")
}
