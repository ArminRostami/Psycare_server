package main

import (
	"fmt"
	"log"
	"psycare/internal/adapters/http"
	"psycare/internal/adapters/postgres"
	app "psycare/internal/application"
	"github.com/go-playground/validator"
)

func main() {
	err := bootstrap()
	if err != nil {
		log.Fatal("err: " + err.Error())
	}
}

func bootstrap() error {
	connStr := "postgres://postgres:example@localhost/postgres?sslmode=disable"
	db, err := postgres.GetDB(connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	store := &postgres.UserStore{DB: db}
	us := &app.UserService{Store: store}
	v := validator.New()
	handler := http.Handler{Us: us,Validate: v}
	handler.Serve()
	return nil
}
