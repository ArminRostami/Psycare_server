package postgres

import (
	"psycare/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	assert := assert.New(t)
	db, err := GetDB(connStr)
	assert.NoError(err)
	repo := &UserStore{DB: db}
	err = repo.AddUser(&domain.User{UserName: "armin", Email: "rostamiarmin@gmail.com", Password: "asdf"})
	assert.NoError(err)
}
func TestAddDesc(t *testing.T) {
	assert := assert.New(t)
	db, err := GetDB(connStr)
	assert.NoError(err)
	repo := &UserStore{DB: db}
	u, err := repo.GetUserWithName("armin")
	assert.NoError(err)
	err = repo.changeDesc(u.ID, "here is ze desclis")
	assert.NoError(err)
}
