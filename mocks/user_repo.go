package mocks

import (
	"psycare/internal/domain"

	"github.com/stretchr/testify/mock"
)

// UserRepoMock is the mock type that satisfies the UserRepo interface
type UserRepoMock struct {
	mock.Mock
	users map[string]domain.User
}

// GetUserWithName impl
func (repo *UserRepoMock) GetUserWithName(username string) (*domain.User, error) {
	user := repo.users[username]
	return &user, nil
}

// AddUser impl
func (repo *UserRepoMock) AddUser(u *domain.User) error {
	if repo.users == nil {
		repo.users = make(map[string]domain.User)
	}
	repo.users[u.UserName] = *u
	return nil
}
