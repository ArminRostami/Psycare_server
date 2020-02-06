package postgres

import (
	"fmt"
	"psycare/internal/domain"

	"github.com/jmoiron/sqlx"
)

// UserStore implements the UserStore interface using postgresql
type UserStore struct {
	DB *sqlx.DB
}

func (us *UserStore) connect(connStr string) error {
	db, err := sqlx.Connect(pgDriver, connStr)
	if err != nil {
		return fmt.Errorf("db connection error: %w", err)
	}
	us.DB = db
	return nil
}

// GetUserWithName _
func (us *UserStore) GetUserWithName(username string) (*domain.User, error) {
	u := &domain.User{}
	err := us.DB.Get(u, "SELECT * FROM users WHERE (username=$1)", username)
	if err != nil {
		return nil, fmt.Errorf("no such user: %w", err)
	}
	return u, nil
}

// AddUser _
func (us *UserStore) AddUser(u *domain.User) error {
	tx, err := us.DB.Beginx()
	if err != nil {
		return fmt.Errorf("transaction start failed: %w", err)
	}
	_, err = tx.NamedExec(`INSERT INTO users (username, email, password, credit) 
				  VALUES (:username,:email,:password,:credit)`, u)
	if err != nil {
		return fmt.Errorf("inserting new user failed: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("inserting new user failed: %w", err)
	}
	return nil
}

func (us *UserStore) changeDesc(id int64, desc string) error {
	tx, err := us.DB.Beginx()
	if err != nil {
		return fmt.Errorf("transaction start failed: %w", err)
	}
	_, err = tx.Exec(`INSERT INTO advisors (id, description) VALUES ($1, $2) 
			 ON CONFLICT(id) DO UPDATE 
			 SET description = $2`, id, desc)
	if err != nil {
		return fmt.Errorf("failed to insert new desc: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to insert new desc: %w", err)
	}
	return nil
}
