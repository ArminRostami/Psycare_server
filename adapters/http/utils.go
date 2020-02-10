package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

type httpError struct {
	errType string
	status  int
	err     error
}

func renderError(w http.ResponseWriter, r *http.Request, e *httpError) {
	log.Printf("%+v", e.err)
	render.Status(r, e.status)
	render.JSON(w, r, render.M{"type": e.errType, "message": e.err.Error(), "status": e.status})
}

func renderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, render.M{"data": data, "status": http.StatusOK})
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
	if !ok {
		return -1, fmt.Errorf("could not cast id to float64")
	}
	return int64(idi), nil
}

func (h *Handler) decodeAndValidate(r *http.Request, dst interface{}) *httpError {
	err := render.DecodeJSON(r.Body, dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request decoding error", err: err}
	}
	err = h.Validate.Struct(dst)
	if err != nil {
		return &httpError{status: http.StatusBadRequest, errType: "request validation error", err: err}
	}
	return nil
}
