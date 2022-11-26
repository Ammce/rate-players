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
	fmt.Println("OVDE SAM ", repoPlayer)
	// createdPlayerResp, err := pri.db.NamedExec(`INSERT INTO players (first_name, last_name, date_of_birth, image_url) VALUES (:first_name, :last_name, :date_of_birth, :image_url)`, &repoPlayer)
	createdPlayerResp, err := pri.db.Exec(`INSERT INTO players (first_name, last_name, date_of_birth, image_url) VALUES ($1, $2, $3, $4)`, "Amel", "Muminovic", "01-06-1994", "sad")
	fmt.Println("U error 0 sam")
	if err != nil {
		// TODO - Throw valid db error
		fmt.Println("U Error 1 sam")
		return nil, err
	}
	playerId, err1 := createdPlayerResp.LastInsertId()

	if err1 != nil {
		fmt.Println("U Error 2 sam")
		// TODO - Throw valid db error
		return nil, err1
	}

	playerToCreate.Id = string(rune(playerId))
	return playerToCreate, nil
}

func NewPlayerRepositoryImpl(db *sqlx.DB) PlayerRepositoryImpl {
	return PlayerRepositoryImpl{
		db: db,
	}
}
