package app

import (
	"fmt"
	"psycare/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	GetUserWithName(username string) (*domain.User, error)
	AddUser(u *domain.User) error
}

type UserService struct {
	Store UserStore
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func (s *UserService) AddUser(u *domain.User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	err = s.Store.AddUser(u)
	if err != nil {
		return fmt.Errorf("error adding user: %w", err)
	}
	return nil
}

func (s *UserService) AuthUser(username, password string) (*domain.User, error) {
	u, err := s.Store.GetUserWithName(username)
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
