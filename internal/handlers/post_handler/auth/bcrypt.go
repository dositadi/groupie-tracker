package auth

import (
	"github.com/dositadi/groupie-tracker/internal/helper"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost                      = 12
	sourceHashpassword string = "Hash password f(n) under auth pkg"
)

func (a *Auth) hashPassword(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		e := helper.WrapError("Password hash error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceHashpassword,
		})
		return nil, e
	}
	return hashedPassword, nil
}

func (a *Auth) compareHashedPassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		e := helper.WrapError("Compare hash and password error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceHashpassword,
		})
		return e
	}
	return nil
}
