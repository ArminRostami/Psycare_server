package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"psycare/internal/adapters/http"
	"psycare/internal/adapters/postgres"
	app "psycare/internal/application"

	"github.com/go-chi/jwtauth"
	"github.com/go-playground/validator"
)

func main() {
	err := bootstrap()
	if err != nil {
		log.Fatal("err: " + err.Error())
	}
}

func bootstrap() error {
	// const connStr = "postgres://postgres:example@localhost/postgres?sslmode=disable"
	connStr := "user=postgres password=example host=localhost port=5432 database=postgres sslmode=disable"
	db, err := postgres.GetDB(connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	userStore := &postgres.UserStore{DB: db}
	advisorStore := &postgres.AdvisorStore{DB: db}
	us := &app.UserService{Store: userStore}
	as := &app.AdvisorService{Store: advisorStore}
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	secret := []byte("VhFJdNDsE9vheq6wTEFga7WhuR4TJ1E8JTPNFaH3e_o")
	auth := jwtauth.New("HS256", secret, nil)
	handler := http.Handler{UserService: us, AdvisorService: as, Validate: v, Auth: auth}
	handler.Serve()
	return nil
}
