package player

// TODO - Add input DTOs to the functions
type PlayerRepository interface {
	CreatePlayer(*Player) (*Player, error)
	GetPlayerById(playerId string) (*Player, error)
	// UpdatePlayer() (*Player, error)
	// GetPlayers() (*[]Player, error)
	// DeletePlayerById() (bool, error)
}
