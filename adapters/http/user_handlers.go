package http

import (
	"net/http"
	"psycare/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/render"
)

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := &domain.User{}
	httpErr := h.decodeAndValidate(r, user)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	err := h.CreateUser(user)
	if err != nil {
		renderError(w, r, &httpError{"user registration error", http.StatusInternalServerError, err})
		return
	}
	user.Password = ""
	renderData(w, r, user)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	req := &struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}{}

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

	roles, err := h.GetRoles(user.ID)
	if err != nil {
		renderError(w, r, &httpError{"failed to get roles", http.StatusInternalServerError, err})
		return
	}
	user.Roles = roles

	_, tokenString, err := h.Auth.Encode(jwt.MapClaims{"id": user.ID})
	if err != nil {
		renderError(w, r, &httpError{"failed to generate token", http.StatusInternalServerError, err})
		return
	}
	http.SetCookie(w, &http.Cookie{Name: "jwt", Value: tokenString})

	renderData(w, r, render.M{"user": user, "cookie": tokenString})
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, httpErr := getIDFromClaims(r)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	u, err := h.GetUserWithID(id)
	if err != nil {
		renderError(w, r, &httpError{"failed to get user", http.StatusInternalServerError, err})
		return
	}
	u.Password = ""
	roles, err := h.GetRoles(id)
	if err != nil {
		renderError(w, r, &httpError{"failed to get roles", http.StatusInternalServerError, err})
		return
	}
	u.Roles = roles
	renderData(w, r, u)
}
