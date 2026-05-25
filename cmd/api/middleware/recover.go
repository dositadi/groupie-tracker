package middleware

import "net/http"

func (m *Middleware) Recover(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				m.handler.ServerErrorHandler(w, r)
			}

		}()
		next.ServeHTTP(w, r)
	})
}
