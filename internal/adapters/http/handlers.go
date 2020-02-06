package http

import (
	"fmt"
	"net/http"
	"psycare/internal/domain"
	app "psycare/internal/domain"

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
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		renderError(w, r, &httpError{"could not get claims", http.StatusInternalServerError, err})
		return
	}
	id, err := getID(claims)
	if err != nil {
		renderError(w, r, &httpError{"could not get id from claims", http.StatusInternalServerError, err})
		return
	}
	a := &domain.Advisor{}

	httpErr := h.decodeAndValidate(r, a)
	if httpErr != nil {
		renderError(w, r, httpErr)
		return
	}
	a.ID = int64(id)
	err = h.CreateAdvisor(a)
	if err != nil {
		renderError(w, r, &httpError{"failed to create advisor", http.StatusInternalServerError, err})
		return
	}
	renderData(w, r, "advisor registered")
}

func getID(claims jwt.MapClaims) (int64, error) {
	id, exists := claims["id"]
	if !exists {
		return -1, fmt.Errorf("claims does not include id")
	}
	idi, ok := id.(float64)
	fmt.Println(idi)
	if !ok {
		return -1, fmt.Errorf("could not cast id to int64")
	}
	return int64(idi), nil
}
