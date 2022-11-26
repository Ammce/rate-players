package repository_impl_player

import (
	"github.com/Ammce/rate-players/src/ports/player"
	"github.com/jmoiron/sqlx"
)

type PlayerRepositoryImpl struct {
	db *sqlx.DB
}

func (pri PlayerRepositoryImpl) CreatePlayer(_ *player.Player) (*player.Player, error) {
	return &player.Player{
		Id: "sjfijsaf",
	}, nil
}

func NewPlayerRepositoryImpl(db *sqlx.DB) PlayerRepositoryImpl {
	return PlayerRepositoryImpl{
		db: db,
	}
}
