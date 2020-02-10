package postgres

import (
	"fmt"
	"psycare/domain"
)

type UserStore struct {
	DB *PDB
}

func (us *UserStore) GetUserWithName(username string) (*domain.User, error) {
	u := &domain.User{}
	err := us.DB.Con.Get(u, "SELECT * FROM users WHERE (username=$1)", username)
	if err != nil {
		return nil, fmt.Errorf("no such user: %w", err)
	}
	return u, nil
}

func (us *UserStore) AddUser(u *domain.User) error {
	return us.DB.namedExec(`INSERT INTO users (username, email, password, credit) 
			  				VALUES (:username,:email,:password,:credit)`, u)
}

func (us *UserStore) changeDesc(id int64, desc string) error {
	return us.DB.exec(`INSERT INTO advisors (id, description) VALUES ($1, $2) 
			 		   ON CONFLICT(id) DO UPDATE 
	 		 		   SET description = $2`, id, desc)
}
