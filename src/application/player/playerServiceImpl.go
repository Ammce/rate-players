package player_service_impl

import (
	"errors"

	player "github.com/Ammce/rate-players/src/ports/player"
)

type PlayerServiceImpl struct {
}

func (PlayerServiceImpl) CreatePlayer(player *player.Player) (*player.Player, error) {
	return player, nil
}

func (PlayerServiceImpl) UpdatePlayer() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func (PlayerServiceImpl) GetPlayerById() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func (PlayerServiceImpl) GetPlayers() (*[]player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func (PlayerServiceImpl) DeletePlayerById() (bool, error) {
	return false, errors.New("Service Method Not Implemented")
}

func NewPlayerServiceImpl() PlayerServiceImpl {
	return PlayerServiceImpl{}
}

func makePlayer() *player.Player {
	return &player.Player{
		Id:          "testId",
		FirstName:   "Amel",
		LastName:    "Muminovic",
		DateOfBirth: "770475473000",
		ImageURL:    "https://static.wikia.nocookie.net/multiversus/images/5/55/Morty_MV.png",
	}
}
