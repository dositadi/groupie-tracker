package middleware

import (
	"github.com/dositadi/groupie-tracker.git/cmd/api/handlers"
	jsonlog "github.com/dositadi/groupie-tracker.git/internal/json_log"
)

type Middleware struct {
	handler handlers.Handler
	logger  jsonlog.Logger
}

func New(handler handlers.Handler, logger jsonlog.Logger) *Middleware {
	return &Middleware{
		handler: handler,
		logger:  logger,
	}
}
