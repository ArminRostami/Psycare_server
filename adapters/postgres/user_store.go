package postgres

import (
	"psycare/domain"

	"github.com/pkg/errors"
)

type UserStore struct {
	DB *PDB
}

func (us *UserStore) GetUserWithName(username string) (*domain.User, error) {
	u := &domain.User{}
	err := us.DB.Con.Get(u, "SELECT * FROM users WHERE (username=$1)", username)
	return u, errors.Wrap(err, "failed to get user")
}

func (us *UserStore) GetUserWithID(id int64) (*domain.User, error) {
	u := &domain.User{}
	err := us.DB.Con.Get(u, "SELECT * FROM users WHERE (id=$1)", id)
	return u, errors.Wrap(err, "failed to get user")
}

func (us *UserStore) CreateUser(u *domain.User) error {
	err := us.DB.NamedExecute(`
	INSERT INTO users (username, email, password) 
	VALUES (:username, :email, :password)`, u)
	return errors.Wrap(err, "failed to create user")
}

func (us *UserStore) Pay(senderID, recieverID, credits int64) error {
	err := us.DB.Execute(`
	UPDATE users AS u SET
	credit = u.credit + u2.credit
	FROM (VALUES ($1::integer, -1*$3),($2::integer, $3)) as u2(id, credit)
	WHERE u.id=u2.id
`, senderID, recieverID, credits)
	return errors.Wrap(err, "payment failed")
}

// func (us *UserStore) changeDesc(id int64, desc string) error {
// 	return us.DB.exec(`INSERT INTO advisors (id, description) VALUES ($1, $2)
// 			 		   ON CONFLICT(id) DO UPDATE
// 	 		 		   SET description = $2`, id, desc)
// }
