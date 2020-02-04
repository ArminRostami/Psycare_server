package http

import (
	"encoding/json"
	"net/http"
	app "psycare/internal/domain"

	"github.com/dgrijalva/jwt-go"
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
	return
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

func (h *Handler) decodeAndValidate(r *http.Request, dst interface{}) *httpError {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request decoding error", err: err}
	}
	err = h.Validate.Struct(dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request validation error", err: err}
	}
	return nil
}
