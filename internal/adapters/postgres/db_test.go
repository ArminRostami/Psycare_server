package postgres

import (
	_ "github.com/jackc/pgx/stdlib"
)

// const pgDriver = "pgx"

// func TestCreateSchema(t *testing.T) {
// 	assert := assert.New(t)
// 	db, err := sqlx.Connect(pgDriver, connStr)
// 	assert.NoError(err)
// 	_, err = db.Exec(DefaultSchema.create)
// 	assert.NoError(err)
// }
// func TestDropSchema(t *testing.T) {
// 	assert := assert.New(t)
// 	db, err := sqlx.Connect(pgDriver, connStr)
// 	assert.NoError(err)
// 	_, err = db.Exec(DefaultSchema.drop)
// 	assert.NoError(err)
// }
