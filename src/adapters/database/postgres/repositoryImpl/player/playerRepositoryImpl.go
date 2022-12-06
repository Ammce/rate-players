package repository_impl_player

import (
	"fmt"

	"github.com/Ammce/rate-players/src/ports/player"
	"github.com/jmoiron/sqlx"
)

// TODO - Move this to separate file

type BaseDB struct {
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	DeletedAt *string `db:"deleted_at"`
}

type RepoPlayer struct {
	BaseDB
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

func RepoPlayerToAppPlayer(repoPlayer RepoPlayer) *player.Player {
	return &player.Player{
		Id:          repoPlayer.Id,
		FirstName:   repoPlayer.FirstName,
		LastName:    repoPlayer.LastName,
		DateOfBirth: repoPlayer.DateOfBirth,
		ImageURL:    repoPlayer.ImageURL,
	}
}

// TODO - Move this to separate file

type PlayerRepositoryImpl struct {
	db *sqlx.DB
}

func (pri PlayerRepositoryImpl) CreatePlayer(playerToCreate *player.Player) (*player.Player, error) {
	repoPlayer := AppPlayerToRepoPlayer(playerToCreate)

	lastInsertId := ""

	createPlayerQuery := `INSERT INTO players (first_name, last_name, date_of_birth, image_url) VALUES ($1, $2, $3, $4) RETURNING id`

	err := pri.db.QueryRow(createPlayerQuery, repoPlayer.FirstName, repoPlayer.LastName, repoPlayer.DateOfBirth, repoPlayer.ImageURL).Scan(&lastInsertId)
	if err != nil {
		// TODO - Throw valid db error
		fmt.Println("U Error 1 sam")
		fmt.Println(err)
		return nil, err
	}

	playerToCreate.Id = lastInsertId
	return playerToCreate, nil
}

func (pri PlayerRepositoryImpl) GetPlayers() ([]*player.Player, error) {
	fmt.Println("GETTING ALL PLAYERS FOR NOW WITHOUT FILTERS")
	return nil, nil
}

func (pri PlayerRepositoryImpl) GetPlayerById(playerId string) (*player.Player, error) {
	fmt.Println("Poziva se")
	repoPlayer := RepoPlayer{}
	err := pri.db.Get(&repoPlayer, "SELECT * FROM players WHERE id=$1", playerId)
	if err != nil {
		// TODO - Throw valid db error
		fmt.Println("U Error 1 sam")
		fmt.Println(err)
		return nil, err
	}

	domainPlayer := RepoPlayerToAppPlayer(repoPlayer)

	return domainPlayer, nil
}

func NewPlayerRepositoryImpl(db *sqlx.DB) PlayerRepositoryImpl {
	return PlayerRepositoryImpl{
		db: db,
	}
}
