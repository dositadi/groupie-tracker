package middleware

import (
	"net/http"

	"github.com/dositadi/groupie-tracker.git/cmd/api/handlers"
)

type Middleware struct {
	handler handlers.Handler
}

func (m *Middleware) Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := recover(); err != nil {
			m.handler.ServerErrorHandler(w, r)
		}
		next.ServeHTTP(w, r)
	})
}
