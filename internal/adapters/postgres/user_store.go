package postgres

import (
	"fmt"
	app "psycare/internal/domain"

	"github.com/jmoiron/sqlx"
)

// UserStore implements the app.UserStore interface using postgresql
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

func (us *UserStore) GetUserWithName(username string) (*app.User, error) {
	u := &app.User{}
	err := us.DB.Get(u, "SELECT * FROM users WHERE (user_name=$1)", username)
	if err != nil {
		return nil, fmt.Errorf("no such user: %w", err)
	}
	return u, nil
}

func (us *UserStore) AddUser(u *app.User) error {
	tx := us.DB.MustBegin()
	tx.NamedExec(`INSERT INTO users (user_name, email, password, credit) 
				  VALUES (:user_name,:email,:password,:credit)`, u)
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("inserting new user failed: %w", err)
	}
	return nil
}

func (us *UserStore) changeDesc(id int64, desc string) error {
	tx := us.DB.MustBegin()
	tx.Exec(`INSERT INTO advisors (id, description) VALUES ($1, $2) 
			 ON CONFLICT(id) DO UPDATE 
			 SET description = $2`, id, desc)
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to insert new desc: %w", err)
	}
	return nil
}
