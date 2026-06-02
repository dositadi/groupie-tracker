package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRL = "Render login page f(n) under authservice pkg"
)

func (a *AuthService) RenderLoginPage() error {
	fs := []string{
		"internal/web/static/auth/login.html",
	}

	temp, err := template.New("login.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRL,
		})
		return e
	}

	data := struct {
		UsernameKey, EmailKey, PasswordKey string
		SignupUrl, LoginUrl string
	}{
		UsernameKey: utils.USERNAME_KEY,
		EmailKey:    utils.EMAIL_KEY,
		PasswordKey: utils.PASSWORD_KEY,
		SignupUrl: utils.REGISTER.String(),
		LoginUrl: utils.LOGIN.String(),
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRL,
		})
		return e
	}
	return nil
}
