package http

import (
	"net/http"
	"psycare/internal/domain"

	"github.com/dgrijalva/jwt-go"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}
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
