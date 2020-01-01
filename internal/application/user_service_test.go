package server

import (
	app "psycare/internal/domain"
	"psycare/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)
	ur := &mocks.UserRepoMock{}
	service := userService{repo: ur}
	err := service.addUser(&app.User{UserName: "armin", Email: "rostamiarmin@gmail.com", Password: "asdfasdf"})
	assert.NoError(err)
	u, err := service.authUser("armin", "asdfasdf")
	assert.NoError(err)
	assert.Equal(u.UserName, "armin")
	assert.Equal(u.Password, "")
	u, err = service.authUser("armin", "asdf")
	assert.Error(err)
	assert.Nil(u)
}
