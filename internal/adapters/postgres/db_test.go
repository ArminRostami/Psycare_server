package postgres

import (
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

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
