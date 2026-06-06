package authpost

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/authservice"
	"github.com/dositadi/groupie-tracker/internal/services/pages"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/dositadi/groupie-tracker/internal/validator"
	"github.com/google/uuid"
)

const (
	sourceReg      = "Register handler f(n) under auth pkg"
	termsUnchecked = "Kindly check the terms."
)

func (a *Auth) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<10) // => 500kb

	if err := r.ParseForm(); err != nil {
		e := helper.WrapError("Body too large", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	username := r.FormValue(utils.USERNAME_KEY)
	password := r.FormValue(utils.PASSWORD_KEY)
	email := r.FormValue(utils.EMAIL_KEY)
	terms := r.FormValue(utils.TERMS_KEY)
	check := true

	authService := authservice.New(w, a.embedded, a.logger)

	errType, err := validator.ValidRegFormValues(username, email, password)
	if err != nil {
		check = false
		switch errType {
		case authservice.EMAIL_ERROR:
			_ = authService.RenderAuthError(errType, err.Error())
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		case authservice.NAME_ERROR:
			_ = authService.RenderAuthError(errType, err.Error())
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		case authservice.PASSWORD_ERROR:
			_ = authService.RenderAuthError(errType, err.Error())
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceReg,
			})
		}
		return
	}

	if !(terms == "true") {
		check = false
		_ = authService.RenderAuthError(errType, termsUnchecked)
		a.logger.PrintError("Terms not checked", map[string]string{
			"Source": sourceReg,
		})
		return
	}

	hashedPassword, err := a.hashPassword([]byte(password))
	if err != nil {
		check = false
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user data.User

	user.Id = uuid.NewString()
	user.Email = email
	user.Username = username
	user.HashedPassword = hashedPassword
	user.Agreed = true

	err = a.usermodel.Insert(user)
	if err != nil {
		check = false
		e := helper.WrapError("User insert error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	defaultPref := data.Preference{
		Id:     uuid.NewString(),
		UserId: user.Id,
		Filter: string(pages.FILTER_BY_ID),
		Sort:   string(pages.ASCENDING_ORDER),
	}

	// (id, userId, filter, sort)
	err = a.preferenceModel.Insert(defaultPref)
	if err != nil {
		check = false
		e := helper.WrapError("Pref insert error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceReg,
		})
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	if check {
		http.Redirect(w, r, utils.LOGIN.String(), http.StatusSeeOther)
	}
}
