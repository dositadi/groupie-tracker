package validator

import (
	"errors"
	"regexp"
)

func ValidRegFormValues(username, email, password string) error {
	switch {
	case username == "":
		return errors.New("Username field cannot be empty")
	case email == "":
		return errors.New("Email field cannot be empty")
	case password == "":
		return errors.New("Password field cannot be empty")
	case email != "":
		if !validateEmail(email) {
			return errors.New("Invalid email address format. Example: example@gmail.com")
		}
	case password != "":
		if len(password) < 8 {
			return errors.New("Password should be at least 8 characters long")
		}
	}
	return nil
}

func ValidateLoginFormValues(email, password string) error {
	switch {
	case email == "":
		return errors.New("Email field cannot be empty")
	case password == "":
		return errors.New("Password field cannot be empty")
	case email != "":
		if !validateEmail(email) {
			return errors.New("Invalid email address format. Example: example@gmail.com")
		}
	case password != "":
		if len(password) < 8 {
			return errors.New("Password should be at least 8 characters long")
		}
	}
	return nil
}

func validateEmail(email string) bool {
	valid, _ := regexp.Match(`[a-zA-Z]+[a-zA-Z0-9-_.]*@[a-z]{8}[.]{1}[a-z]{3}`, []byte(email))

	return valid
}
