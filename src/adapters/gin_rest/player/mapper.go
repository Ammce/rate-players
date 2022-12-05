package player_http

import "github.com/Ammce/rate-players/src/ports/player"

type CreatePlayerInput struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	ImageURL    string `json:"image_url"`
}

func CreatePlayerInputToPlayerEntity(createPlayerInput CreatePlayerInput) player.Player {
	return player.Player{
		FirstName:   createPlayerInput.FirstName,
		LastName:    createPlayerInput.LastName,
		ImageURL:    createPlayerInput.ImageURL,
		DateOfBirth: createPlayerInput.DateOfBirth,
	}
}
