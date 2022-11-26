package player

// TODO - Add input DTOs to the functions
type PlayerRepository interface {
	CreatePlayer(*Player) (*Player, error)
	// UpdatePlayer() (*Player, error)
	// GetPlayerById() (*Player, error)
	// GetPlayers() (*[]Player, error)
	// DeletePlayerById() (bool, error)
}
