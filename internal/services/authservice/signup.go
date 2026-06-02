package authservice

import (
	"html/template"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRS = "Render signup page f(n) under authservice"
)

func (a *AuthService) RenderSignupPage() error {
	fs := []string{
		"internal/web/static/auth/signup.html",
	}

	temp, err := template.New("signup.html").ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}

	data := struct {
		SignupUrl, LoginUrl                string
		PrivacyUrl, TermUrl                string
		UsernameKey, EmailKey, PasswordKey string
	}{
		SignupUrl:   utils.REGISTER.String(),
		LoginUrl:    utils.LOGIN.String(),
		UsernameKey: utils.USERNAME_KEY,
		EmailKey:    utils.EMAIL_KEY,
		PasswordKey: utils.PASSWORD_KEY,
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Error executing template", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}
	return nil
}
