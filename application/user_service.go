package app

import (
	"psycare/domain"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	GetUserWithName(username string) (*domain.User, error)
	GetUserWithID(id int64) (*domain.User, error)
	CreateUser(u *domain.User) error
	Pay(senderID, recieverID, credits int64) error
}

type UserService struct {
	UserStore
	RoleStore
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.Wrap(err, "failed to hash password")
	}
	return string(hash), nil
}

func (us *UserService) CreateUser(u *domain.User) error {
	hash, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	err = us.UserStore.CreateUser(u)
	if err != nil {
		return errors.WithMessage(err, "failed to add user")
	}
	return nil
}

func (us *UserService) AuthUser(username, password string) (*domain.User, error) {
	u, err := us.UserStore.GetUserWithName(username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, errors.Wrap(err, "password mismatch")
	}
	u.Password = ""
	return u, nil
}

func (us *UserService) GetUserWithID(id int64) (*domain.User, error) {
	return us.UserStore.GetUserWithID(id)
}
