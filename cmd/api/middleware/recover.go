package middleware

import "net/http"

func (m *Middleware) Recover(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				e, ok := err.(string)
				if !ok {
					e = "A panic occurred within the program handlers."
				}
				m.logger.PrintError(e, map[string]string{
					"Source": "Recover middleware",
				})
				m.handler.ServerErrorHandler(w, r)
			}

		}()
		next.ServeHTTP(w, r)
	})
}
