package domain

type Advisor struct {
	ID          int64  `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name" validate:"required"`
	LastName    string `db:"last_name" json:"last_name" validate:"required"`
	Description string `db:"description" json:"description" validate:"required"`
	Verified    bool   `db:"verified" json:"verified,omitempty"`
	HourlyFee   int64  `db:"hourly_fee" json:"hourly_fee,omitempty"`
}
