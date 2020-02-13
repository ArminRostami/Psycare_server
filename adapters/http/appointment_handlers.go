package http

import (
	"net/http"
	"psycare/domain"
)

func (h *Handler) bookAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appt := &domain.Appointment{}
	appt.UserID = id
	httpErr = h.decodeAndValidate(r, appt)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	// TODO: assert that appointment can be booked

	cost, err := h.CalculateCost(appt)
	if err != nil {
		renderError(w, r, &httpError{"cannot calculate cost for appointment", http.StatusInternalServerError, err})
		return
	}

	err = h.Pay(appt.UserID, appt.AdvisorID, cost)
	if err != nil {
		renderError(w, r, &httpError{"payment failed", http.StatusInternalServerError, err})
		return
	}

	err = h.CreateAppointment(appt)
	if err != nil {
		renderError(w, r, &httpError{"failed to add appointment", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, appt)
}

func (h *Handler) appointmentsHandler(w http.ResponseWriter, r *http.Request, forUser bool) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	appts, err := h.GetAppointments(id, forUser)
	if err != nil {
		renderError(w, r, &httpError{"failed to get appointments", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, appts)
}

func (h *Handler) getUserAppointments(w http.ResponseWriter, r *http.Request) {
	h.appointmentsHandler(w, r, true)
}

func (h *Handler) getAdvisorAppointments(w http.ResponseWriter, r *http.Request) {
	h.appointmentsHandler(w, r, false)
}

func (h *Handler) rateAppointment(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	rating := &domain.Rating{UserID: id}
	httpErr = h.decodeAndValidate(r, rating)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	err := h.AddRating(rating)
	if err != nil {
		renderError(w, r, &httpError{"failed to add rating", http.StatusInternalServerError, err})
		return
	}

	renderData(w, r, rating)

}
