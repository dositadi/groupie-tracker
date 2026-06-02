package auth

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/authservice"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/dositadi/groupie-tracker/internal/validator"
	"github.com/google/uuid"
)

const (
	sourceReg     = "Register handler f(n) under auth pkg"
	usernameEmpty = "Username field cannot be empty"
	passwordEmpty = "Password field cannot be empty/Password less than 8 characters"
	emailEmpty    = "Email field cannot be empty/Invalid Email"
)

func (a *Auth) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<10) // => 500kb

	if err := r.ParseForm(); err != nil {
		e := helper.WrapError("Body too large", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	username := r.FormValue(utils.USERNAME_KEY)
	password := r.FormValue(utils.PASSWORD_KEY)
	email := r.FormValue(utils.EMAIL_KEY)
	check := true

	authService := authservice.New(w, a.embedded, a.logger)

	errType, err := validator.ValidRegFormValues(username, email, password)
	if err != nil {
		check = false
		switch errType {
		case authservice.EMAIL_ERROR:
			authService.RenderAuthError(errType, emailEmpty)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		case authservice.NAME_ERROR:
			authService.RenderAuthError(errType, usernameEmpty)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		case authservice.PASSWORD_ERROR:
			authService.RenderAuthError(errType, passwordEmpty)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		}
	}

	hashedPassword, err := a.hashPassword([]byte(password))
	if err != nil {
		check = false
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var user data.User

	user.Id = uuid.NewString()
	user.Email = email
	user.Username = username
	user.HashedPassword = hashedPassword

	err = a.usermodel.Insert(user)
	if err != nil {
		check = false
		e := helper.WrapError("User insert error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	if check {
		http.Redirect(w, r, utils.LOGIN.String(), http.StatusSeeOther)
	}
}
