package mocks

import (
	"psycare/internal/domain"

	"github.com/stretchr/testify/mock"
)

// UserStoreMock is the mock type that satisfies the UserRepo interface
type UserStoreMock struct {
	mock.Mock
	users map[string]domain.User
}

// GetUserWithName impl
func (repo *UserStoreMock) GetUserWithName(username string) (*domain.User, error) {
	user := repo.users[username]
	return &user, nil
}

// AddUser impl
func (repo *UserStoreMock) AddUser(u *domain.User) error {
	if repo.users == nil {
		repo.users = make(map[string]domain.User)
	}
	repo.users[u.UserName] = *u
	return nil
}
