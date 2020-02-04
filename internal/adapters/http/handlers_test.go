package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	app "psycare/internal/application"
	"psycare/internal/domain"
	"psycare/mocks"
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	us := &app.UserService{Store: &mocks.UserRepoMock{}}
	v := validator.New()
	h := &Handler{Validate: v, Us: us}
	w := httptest.NewRecorder()
	u, err := json.Marshal(domain.User{UserName: "armin", Password: "asdfasdf", Email: "rostamiarmin@gmail.com"})
	assert.NoError(err)
	r := httptest.NewRequest("POST", "http://example.com", bytes.NewReader(u))
	h.addUser(w, r)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	t.Log(string(body))
}
