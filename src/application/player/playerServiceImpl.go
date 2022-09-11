package player_service_impl

import (
	"errors"

	player "github.com/Ammce/rate-players/src/ports/player"
)

type PlayerServiceImpl struct {
}

func CreatePlayer() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func UpdatePlayer() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func GetPlayerById() (*player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func GetPlayers() (*[]player.Player, error) {
	return nil, errors.New("Service Method Not Implemented")
}

func DeletePlayerById() (bool, error) {
	return false, errors.New("Service Method Not Implemented")
}

func makePlayer() *player.Player {
	return &player.Player{
		Id:             "testId",
		FirstName:      "Amel",
		LasttName:      "Muminovic",
		DateOfBirth:    770475473000,
		Nationality:    "SRB",
		ProfilePicture: "https://static.wikia.nocookie.net/multiversus/images/5/55/Morty_MV.png",
	}
}
