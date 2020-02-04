package http

import (
	"encoding/json"
	"net/http"
	app "psycare/internal/domain"
)

func (h *Handler) addUser(w http.ResponseWriter, r *http.Request) {
	user := &app.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "request decoding error", err)
		return
	}
	err = h.Validate.Struct(user)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "request validation error", err)
		return
	}
	err = h.Us.AddUser(user)
	if err != nil {
		renderError(w, r, http.StatusInternalServerError, "user registration error", err)
		return
	}
	renderData(w, r, "user registered")
	return
}
