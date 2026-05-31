package auth

import (
	"fmt"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/dositadi/groupie-tracker/internal/validator"
	"github.com/google/uuid"
)

const (
	sourceReg     = "Register handler f(n) under auth pkg"
	usernameEmpty = "Username field cannot be empty"
	passwordEmpty = "Password field cannot be empty"
	emailEmpty    = "Email field cannot be empty"
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

	err := validator.ValidRegFormValues(username, email, password)
	if err != nil {
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	hashedPassword, err := a.hashPassword([]byte(password))
	if err != nil {
		e := fmt.Errorf(err.Error())
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	var user data.User

	user.Id = uuid.NewString()
	user.Email = email
	user.Username = username
	user.HashedPassword = hashedPassword

	err = a.usermodel.Insert(user)
	if err != nil {
		e := helper.WrapError("User insert error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, utils.LOGIN.String(), http.StatusSeeOther)
}
