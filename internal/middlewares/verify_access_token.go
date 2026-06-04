package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Middleware) VerifyAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// Step 1: get the token string form either header or cookie
			tokenString := m.getToken(r)
			if tokenString == "" {
				// Direct to the page for expired session page
				e := fmt.Errorf("1. Unauthorized access %s", tokenString)
				m.logger.PrintError(e.Error(), map[string]string{
					"Source": "Verify access token f(n) under middleware pkg",
				})
				http.Error(w, e.Error(), http.StatusUnauthorized)
				return
			}

			// Parse the claim using the address of the variable below
			var active data.ActiveUser
			token, err := jwt.ParseWithClaims(tokenString, &active, func(t *jwt.Token) (any, error) {
				jwtKey := os.Getenv("JWTKEY")
				return []byte(jwtKey), nil
			})
			if err != nil {
				// Direct to the page for expired session page
				e := helper.WrapError("2. Unauthorized access", err)
				logger.PrintError(e.Error()+" "+tokenString, map[string]string{
					"Source": "Verify access token f(n) under middleware pkg",
				})
				http.Error(w, e.Error(), http.StatusUnauthorized)
				return
			}

			// Check that the token is valid and is of the active user type
			if _, ok := token.Claims.(*data.ActiveUser); !ok && !token.Valid {
				e := helper.WrapError("3. Unauthorized access", err)
				logger.PrintError(e.Error(), map[string]string{
					"Source": "Verify access token f(n) under middleware pkg",
				})
				http.Error(w, e.Error(), http.StatusUnauthorized)
				return
			}

			cxt := context.WithValue(r.Context(), utils.USER_ID_KEY, active.Id)
			fmt.Println(active.Id)
			next.ServeHTTP(w, r.WithContext(cxt))
		},
	)
}

func (m *Middleware) getToken(r *http.Request) string {
	cookie, err := r.Cookie(utils.ACCESS_TOKEN_KEY)
	if err == nil {
		return cookie.Value
	}
	return ""
}
