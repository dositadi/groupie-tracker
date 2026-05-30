package models

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	favoritemodel "github.com/dositadi/groupie-tracker/internal/models/favorite_model"
	searchmodel "github.com/dositadi/groupie-tracker/internal/models/search_model"
	usermodel "github.com/dositadi/groupie-tracker/internal/models/user_model"
	"github.com/jackc/pgx/v5"
)

type Models struct {
	UserModel     usermodel.UserModel
	FavoriteModel favoritemodel.FavoriteModel
	SearchModel   searchmodel.SearchModel
}

func New(db *pgx.Conn, logger jsonlog.Logger) *Models {
	return &Models{
		UserModel:     *usermodel.New(db, logger),
		FavoriteModel: *favoritemodel.New(db, logger),
		SearchModel:   *searchmodel.New(db, logger),
	}
}
