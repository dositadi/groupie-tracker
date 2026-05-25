package middleware

import (
	"github.com/dositadi/groupie-tracker.git/cmd/api/handlers"
)

type Middleware struct {
	handler handlers.Handler
}

func New(handler handlers.Handler) *Middleware {
	return &Middleware{
		handler: handler,
	}
}
