package server

import (
	"fmt"
	app "psycare/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

// UserService extends UserRepo to add more functionality
type UserService struct {
	Repo app.UserRepo
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

// AddUser adds the user to repository
func (s *UserService) AddUser(u *app.User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	err = s.Repo.AddUser(u)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}

func (s *UserService) authUser(username, password string) (*app.User, error) {
	u, err := s.Repo.GetUserWithName(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password mismatch: %w", err)
	}
	u.Password = ""
	return u, nil
}
