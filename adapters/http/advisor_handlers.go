package http

import (
	"net/http"
	"psycare/domain"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) createAdvisor(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	a := &domain.Advisor{}
	httpErr = h.decodeAndValidate(r, a)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	a.ID = id
	err := h.CreateAdvisor(a)
	if err != nil {
		renderError(w, r, &httpError{"failed to create advisor", http.StatusInternalServerError, err})
		return
	}

	renderData(w, r, a)
}

func (h *Handler) getAdvisors(w http.ResponseWriter, r *http.Request) {

	var limit, offset int
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	if limitStr != "" {
		lim, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			renderError(w, r, &httpError{"param parse error", http.StatusBadRequest, err})
			return
		}
		limit = int(lim)
	} else {
		limit = 20
	}
	if offsetStr != "" {
		off, err := strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			renderError(w, r, &httpError{"param parse error", http.StatusBadRequest, err})
			return
		}
		offset = int(off)
	}

	advisors, err := h.GetAdvisors(true, limit, offset)
	if err != nil {
		renderError(w, r, &httpError{"resource fetching error", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, advisors)
}

func (h *Handler) addSchedule(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	sch := &domain.Schedule{}
	httpErr = h.decodeAndValidate(r, sch)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	sch.AdvisorID = id
	err := h.AddSchedule(sch)
	if err != nil {
		renderError(w, r, &httpError{"failed to add schedule", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, sch)
}

func (h *Handler) getAvgRating(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "adv_id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		renderError(w, r, &httpError{"no advisor id in url", http.StatusBadRequest, err})
		return
	}

	rating, err := h.GetAvgRating(id64)
	if err != nil {
		renderError(w, r, &httpError{"failed to get avg rating", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, rating)
}

func (h *Handler) getAdvisor(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "adv_id")
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		renderError(w, r, &httpError{"failed to get id from url", http.StatusBadRequest, err})
		return
	}

	adv, err := h.GetAdvisorWithID(id64)
	if err != nil {
		renderError(w, r, &httpError{"failed to get advisor from service", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, adv)
}
