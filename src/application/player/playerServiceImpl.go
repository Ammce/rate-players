package player_service_impl

import (
	"errors"

	player "github.com/Ammce/rate-players/src/ports/player"
)

type PlayerServiceImpl struct {
	playerRepository player.PlayerRepository
}

func (psi PlayerServiceImpl) CreatePlayer(player *player.Player) (*player.Player, error) {
	createdPlayer, err := psi.playerRepository.CreatePlayer(player)
	if err != nil {
		return nil, err
	}
	return createdPlayer, nil
}

func (PlayerServiceImpl) UpdatePlayer() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func (psi PlayerServiceImpl) GetPlayerById(playerId string) (*player.Player, error) {
	playerById, err := psi.playerRepository.GetPlayerById(playerId)
	if err != nil {
		return nil, err
	}
	return playerById, nil
}

func (PlayerServiceImpl) GetPlayers() ([]*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func (PlayerServiceImpl) DeletePlayerById() (bool, error) {
	return false, errors.New("Service Method Not Implemented")
}

func NewPlayerServiceImpl(playerRepository player.PlayerRepository) PlayerServiceImpl {
	return PlayerServiceImpl{
		playerRepository: playerRepository,
	}
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
