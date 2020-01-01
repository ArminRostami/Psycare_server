package server

import (
	"fmt"
	app "psycare/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo app.UserRepo
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (s *userService) addUser(u *app.User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	err = s.repo.AddUser(u)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}

func (s *userService) authUser(username, password string) (*app.User, error) {
	u, err := s.repo.GetUserWithName(username)
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
