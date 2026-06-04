package handlers

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	deletehandler "github.com/dositadi/groupie-tracker/internal/handlers/delete_handler"
	gethandler "github.com/dositadi/groupie-tracker/internal/handlers/get_handler"
	posthandler "github.com/dositadi/groupie-tracker/internal/handlers/post_handler"
	"github.com/dositadi/groupie-tracker/internal/handlers/post_handler/pages"
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
	client    artistapi.ArtistInfo
	Get       gethandler.Get
	Post      posthandler.Post
	Delete    deletehandler.Delete
	embedded  groupietracker.Embedded
}

func New(logger jsonlog.Logger, userModel UserModel, favoriteModel pages.FavoriteModel, client artistapi.ArtistInfo, embedded groupietracker.Embedded) *Handler {
	return &Handler{
		logger:    logger,
		userModel: userModel,
		Get:       *gethandler.New(userModel, client, logger, embedded),
		Post:      *posthandler.New(userModel, favoriteModel, client, logger, embedded),
		Delete:    *deletehandler.New(userModel, client, logger),
	}
}

func (h *Handler) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("An internal server error occurred."))
}
