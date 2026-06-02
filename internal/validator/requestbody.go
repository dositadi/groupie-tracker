package validator

import (
	"errors"
	"regexp"

	"github.com/dositadi/groupie-tracker/internal/services/authservice"
)

func ValidRegFormValues(username, email, password string) (authservice.Error, error) {
	switch {
	case username == "":
		return authservice.NAME_ERROR, errors.New("Username field cannot be empty")
	case email == "":
		return authservice.EMAIL_ERROR, errors.New("Email field cannot be empty")
	case password == "":
		return authservice.PASSWORD_ERROR, errors.New("Password field cannot be empty")
	case email != "":
		if !validateEmail(email) {
			return authservice.EMAIL_ERROR, errors.New("Invalid email address format. Example: example@gmail.com")
		}
	case password != "":
		if len(password) < 8 {
			return authservice.PASSWORD_ERROR, errors.New("Password should be at least 8 characters long")
		}
	}
	return authservice.Error(""), nil
}

func ValidateLoginFormValues(email, password string) (authservice.Error, error) {
	switch {
	case email == "":
		return authservice.EMAIL_ERROR, errors.New("Email field cannot be empty")
	case password == "":
		return authservice.PASSWORD_ERROR, errors.New("Password field cannot be empty")
	case email != "":
		if !validateEmail(email) {
			return authservice.EMAIL_ERROR, errors.New("Invalid email address format. Example: you@example.com")
		}
	case password != "":
		if len(password) < 8 {
			return authservice.PASSWORD_ERROR, errors.New("Password should be at least 8 characters long")
		}
	}
	return authservice.Error(""), nil
}

func validateEmail(email string) bool {
	valid, _ := regexp.Match(`[a-zA-Z]+[a-zA-Z0-9-_.]*@[a-z]{8}[.]{1}[a-z]{3}`, []byte(email))

	return valid
}
