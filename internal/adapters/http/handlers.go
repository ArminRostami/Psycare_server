package http

import (
	"fmt"
	"net/http"
	"psycare/internal/domain"
	app "psycare/internal/domain"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &app.User{}
	httpErr := h.decodeAndValidate(r, user)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	err := h.AddUser(user)
	if err != nil {
		renderError(w, r, &httpError{"user registration error", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, "user registered")
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	req := &request{}

	httpErr := h.decodeAndValidate(r, req)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}

	user, err := h.AuthUser(req.Username, req.Password)
	if err != nil {
		renderError(w, r, &httpError{"auth failed", http.StatusForbidden, err})
		return
	}

	_, tokenString, err := h.Auth.Encode(jwt.MapClaims{"id": user.ID})
	if err != nil {
		renderError(w, r, &httpError{"failed to generate token", http.StatusInternalServerError, err})
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenString})
}

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
	a.ID = int64(id)
	err := h.CreateAdvisor(a)
	if err != nil {
		renderError(w, r, &httpError{"failed to create advisor", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, "advisor registered")
}

func getIDFromClaims(r *http.Request) (int64, *httpError) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return -1, &httpError{"could not get claims", http.StatusInternalServerError, err}
	}
	id, err := getID(claims)
	if err != nil {
		return -1, &httpError{"could not get id from claims", http.StatusInternalServerError, err}
	}
	return id, nil
}

func getID(claims jwt.MapClaims) (int64, error) {
	id, exists := claims["id"]
	if !exists {
		return -1, fmt.Errorf("claims does not include id")
	}
	idi, ok := id.(float64)
	fmt.Println(idi)
	if !ok {
		return -1, fmt.Errorf("could not cast id to float64")
	}
	return int64(idi), nil
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
