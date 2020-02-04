package app

import (
	"psycare/internal/domain"
	"psycare/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuth(t *testing.T) {
	assert := assert.New(t)
	ur := &mocks.UserStoreMock{}
	service := UserService{Store: ur}
	err := service.AddUser(&domain.User{UserName: "armin", Email: "rostamiarmin@gmail.com", Password: "asdfasdf"})
	assert.NoError(err)
	u, err := service.AuthUser("armin", "asdfasdf")
	assert.NoError(err)
	assert.Equal(u.UserName, "armin")
	assert.Equal(u.Password, "")
	u, err = service.AuthUser("armin", "asdf")
	assert.Error(err)
	assert.Nil(u)
}
