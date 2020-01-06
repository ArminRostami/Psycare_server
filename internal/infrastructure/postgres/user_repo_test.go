package postgres

import (
	app "psycare/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	assert := assert.New(t)
	db, err := getDB(connStr)
	assert.NoError(err)
	repo := &UserRepo{db: db}
	err = repo.addUser(&app.User{UserName: "armin", Email: "rostamiarmin@gmail.com", Password: "asdf"})
	assert.NoError(err)
}
func TestAddDesc(t *testing.T) {
	assert := assert.New(t)
	db, err := getDB(connStr)
	assert.NoError(err)
	repo := &UserRepo{db: db}
	u, err := repo.getUserWithName("armin")
	assert.NoError(err)
	err = repo.changeDesc(u.ID, "here is ze desclis")
	assert.NoError(err)
}
