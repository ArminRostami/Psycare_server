package postgres

import (
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const pgDriver = "pgx"
const connStr = "user=postgres password=example host=localhost port=5432 database=psycare sslmode=disable"

func TestCreateSchema(t *testing.T) {

	assert := assert.New(t)
	db, err := sqlx.Connect(pgDriver, connStr)
	assert.NoError(err)
	_, err = db.Exec(DefaultSchema.create)
	assert.NoError(err)
}
func TestDropSchema(t *testing.T) {
	assert := assert.New(t)
	db, err := sqlx.Connect(pgDriver, connStr)
	assert.NoError(err)
	_, err = db.Exec(DefaultSchema.drop)
	assert.NoError(err)
}
