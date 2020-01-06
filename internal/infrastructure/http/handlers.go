package http

import (
	"encoding/json"
	"net/http"
	app "psycare/internal/domain"
)

func (h *handler) addUser(w http.ResponseWriter, r *http.Request) {
	user := &app.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "request decoding error", err)
		return
	}
	err = h.validate.Struct(user)
	if err != nil {
		renderError(w, r, http.StatusBadRequest, "request validation error", err)
		return
	}
	err = h.us.AddUser(user)
	if err != nil {
		renderError(w, r, http.StatusInternalServerError, "user registration error", err)
		return
	}
	renderData(w, r, "user registered")
	return
}
