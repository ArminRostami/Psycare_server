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
	connStr := "user=postgres password=example host=localhost port=5432 database=postgres sslmode=disable"
	pdb, err := postgres.Connect(connStr)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	apptStore := &postgres.AppointmentStore{DB: pdb}
	userStore := &postgres.UserStore{DB: pdb}
	advisorStore := &postgres.AdvisorStore{DB: pdb}
	us := &app.UserService{Store: userStore}
	as := &app.AdvisorService{Store: advisorStore}
	aps := &app.AppointmentService{Store: apptStore}
	srv := &app.Services{UserService: us, AdvisorService: as, AppointmentService: aps}
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
	handler := http.Handler{Services: srv, Validate: v, Auth: auth}
	handler.Serve()
	return nil
}
