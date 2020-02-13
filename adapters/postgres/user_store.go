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
	if err != nil {
		return nil, errors.Wrap(err, "no such user")
	}
	return u, nil
}

func (us *UserStore) GetUserWithID(id int64) (*domain.User, error) {
	u := &domain.User{}
	err := us.DB.Con.Get(u, "SELECT * FROM users WHERE (id=$1)", id)
	if err != nil {
		return nil, errors.Wrap(err, "no such user")
	}
	return u, nil
}

func (us *UserStore) AddUser(u *domain.User) error {
	return us.DB.namedExec(`
	INSERT INTO users (username, email, password) 
	VALUES (:username, :email, :password)`, u)
}

func (us *UserStore) Pay(senderID, recieverID, credits int64) error {
	senderCredits := new(int64)
	err := us.DB.Con.Get(senderCredits, `
	SELECT credit FROM users WHERE id=$1`, senderID)
	if err != nil {
		return errors.Wrap(err, "cannot get sender credits")
	}

	if *senderCredits < credits {
		return errors.New("sender does not have enough credit")
	}

	return us.DB.exec(`
	UPDATE users AS u SET
	credit = u.credit + u2.credit
	FROM (VALUES ($1::integer, -1*$3),($2::integer, $3)) as u2(id, credit)
	WHERE u.id=u2.id
`, senderID, recieverID, credits)
}

// func (us *UserStore) changeDesc(id int64, desc string) error {
// 	return us.DB.exec(`INSERT INTO advisors (id, description) VALUES ($1, $2)
// 			 		   ON CONFLICT(id) DO UPDATE
// 	 		 		   SET description = $2`, id, desc)
// }
