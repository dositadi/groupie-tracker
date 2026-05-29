package handlers

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}

type Handler struct {
	logger    jsonlog.Logger
	userModel UserModel
}

func New(logger jsonlog.Logger, userModel UserModel) *Handler {
	return &Handler{
		logger:    logger,
		userModel: userModel,
	}
}

func (h *Handler) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("An internal server error occurred."))
}
