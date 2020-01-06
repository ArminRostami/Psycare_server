package app

// User is the struct that holds user data
type User struct {
	UserName string  `db:"user_name" json:"user_name" validate:"required"`
	Email    string  `db:"email" json:"email" validate:"required,email"`
	Password string  `db:"password" json:"password" validate:"required"`
	ID       int64   `db:"id"`
	Credit   float64 `db:"credit"`
}

// UserRepo interface defines methods for processing user data
type UserRepo interface {
	GetUserWithName(username string) (*User, error)
	AddUser(u *User) error
}
