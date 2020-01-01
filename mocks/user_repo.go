package mocks

import (
	"github.com/stretchr/testify/mock"

	app "psycare/internal/domain"
)

// UserRepoMock is the mock type that satisfies the UserRepo interface
type UserRepoMock struct {
	mock.Mock
	users map[string]app.User
}

// GetUserWithName impl
func (repo *UserRepoMock) GetUserWithName(username string) (*app.User, error) {
	user := repo.users[username]
	return &user, nil
}

// AddUser impl
func (repo *UserRepoMock) AddUser(u *app.User) error {
	if repo.users == nil {
		repo.users = make(map[string]app.User)
	}
	repo.users[u.UserName] = *u
	return nil
}
