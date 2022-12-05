package repository_impl_player

import (
	"fmt"

	"github.com/Ammce/rate-players/src/ports/player"
	"github.com/jmoiron/sqlx"
)

// TODO - Move this to separate file

type RepoPlayer struct {
	Id          string `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	DateOfBirth string `db:"date_of_birth"`
	ImageURL    string `db:"image_url"`
}

func AppPlayerToRepoPlayer(appPlayer *player.Player) RepoPlayer {
	return RepoPlayer{
		FirstName:   appPlayer.FirstName,
		LastName:    appPlayer.LastName,
		DateOfBirth: appPlayer.DateOfBirth,
		ImageURL:    appPlayer.ImageURL,
	}
}

// TODO - Move this to separate file

type PlayerRepositoryImpl struct {
	db *sqlx.DB
}

func (pri PlayerRepositoryImpl) CreatePlayer(playerToCreate *player.Player) (*player.Player, error) {
	repoPlayer := AppPlayerToRepoPlayer(playerToCreate)

	lastInsertId := ""

	err := pri.db.QueryRow(`INSERT INTO players (first_name, last_name, date_of_birth, image_url) VALUES ($1, $2, $3, $4) RETURNING id`, repoPlayer.FirstName, repoPlayer.LastName, repoPlayer.DateOfBirth, repoPlayer.ImageURL).Scan(&lastInsertId)
	if err != nil {
		// TODO - Throw valid db error
		fmt.Println("U Error 1 sam")
		fmt.Println(err)
		return nil, err
	}

	playerToCreate.Id = lastInsertId
	return playerToCreate, nil
}

func NewPlayerRepositoryImpl(db *sqlx.DB) PlayerRepositoryImpl {
	return PlayerRepositoryImpl{
		db: db,
	}
}
