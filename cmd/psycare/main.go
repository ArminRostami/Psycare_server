package main

import (
	"fmt"
	"log"
	"psycare/adapters/http"
	"psycare/adapters/postgres"
	app "psycare/application"
	"reflect"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var required = []string{
	"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME", "APP_PORT", "JWT_SECRET",
}

func main() {
	err := bootstrap()
	if err != nil {
		log.Fatal("error: " + err.Error())
	}
}

func bootstrap() error {
	env, err := getEnvMap(required)
	if err != nil {
		return err
	}

	pdb, err := postgres.Connect(getConnString(env))
	if err != nil {
		return errors.Wrap(err, "failed to bootstrap application")
	}

	_ = pdb.Execute(postgres.DefaultSchema.Create)

	roleStore := &postgres.RoleStore{DB: pdb}
	userStore := &postgres.UserStore{DB: pdb}
	advisorStore := &postgres.AdvisorStore{DB: pdb}
	apptStore := &postgres.AppointmentStore{DB: pdb}

	uss := &app.UserService{UserStore: userStore, RoleStore: roleStore}
	ads := &app.AdvisorService{AdvisorStore: advisorStore, RoleStore: roleStore}
	aps := &app.AppointmentService{AppointmentStore: apptStore, AdvisorStore: advisorStore, UserStore: userStore}

	srv := &app.Services{UserService: uss, AdvisorService: ads, AppointmentService: aps}

	v := getValidator()

	secret := []byte(env["JWT_SECRET"])
	auth := jwtauth.New("HS256", secret, nil)

	handler := http.Handler{Services: srv, Validate: v, Auth: auth}
	err = handler.Serve(env["APP_PORT"])
	if err != nil {
		return err
	}

	return nil
}

func getConnString(env map[string]string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s database=%s sslmode=disable",
		env["DB_USER"], env["DB_PASS"], env["DB_HOST"], env["DB_PORT"], env["DB_NAME"])
}

func getValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
	return v
}

func getEnvMap(keys []string) (map[string]string, error) {
	env, err := godotenv.Read(".env")
	if err != nil {
		return nil, errors.Wrap(err, "could not read env file")
	}
	for _, key := range keys {
		if _, ok := env[key]; !ok {
			return nil, errors.Errorf(`key "%s" is missing from .env file.`, key)
		}

	}
	return env, nil
}
