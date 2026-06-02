package auth

import (
	"errors"
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/authservice"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/dositadi/groupie-tracker/internal/validator"
)

const (
	sourceLogin = "Login handler f(n) under auth pkg"
)

var (
	INVALID      error = errors.New("Invalid credentials")
	SERVER_ERROR       = errors.New("Something wrong just occured")
)

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	if err := r.ParseForm(); err != nil {
		e := helper.WrapError("Body too large", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceLogin,
		})
		http.Error(w, e.Error(), http.StatusBadRequest)
	}

	email := r.FormValue(utils.EMAIL_KEY)
	password := r.FormValue(utils.PASSWORD_KEY)
	check := true

	authService := authservice.New(w, a.embedded, a.logger)

	errType, err := validator.ValidateLoginFormValues(email, password)
	if err != nil {
		check = false
		switch errType {
		case authservice.EMAIL_ERROR:
			_ = authService.RenderAuthError(errType, emailEmpty)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceLogin,
			})
		case authservice.PASSWORD_ERROR:
			_ = authService.RenderAuthError(errType, emailEmpty)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceLogin,
			})
		}
	}

	user, err := a.usermodel.GetWithEmail(email)
	if err != nil {
		check = false
		a.logger.PrintError(INVALID.Error(), map[string]string{
			"Source": sourceLogin,
			"Email":  email,
		})

		http.Error(w, INVALID.Error(), http.StatusUnauthorized)
	}

	err = a.compareHashedPassword(user.HashedPassword, []byte(password))
	if err != nil {
		check = false
		e := helper.WrapError("User fetch error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source":   sourceLogin,
			"Password": password,
		})
		http.Error(w, INVALID.Error(), http.StatusUnauthorized)
	}

	var activeUser data.ActiveUser
	activeUser.Email = user.Email
	activeUser.Username = user.Username

	token, err := a.generateJWT(activeUser)
	if err != nil {
		check = true
		e := helper.WrapError("JWT generation error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source":   sourceLogin,
			"Password": password,
		})
		http.Error(w, SERVER_ERROR.Error(), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     utils.ACCESS_TOKEN_KEY,
		Value:    string(token),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Domain:   "http://localhost:8080",
	})

	if check {
		http.Redirect(w, r, utils.HOME.String(), http.StatusSeeOther)
	}
}
