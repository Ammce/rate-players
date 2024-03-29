package player

// TODO - Add input DTOs to the functions
type PlayerService interface {
	CreatePlayer(*Player) (*Player, error)
	UpdatePlayer() (*Player, error)
	GetPlayerById(playerId string) (*Player, error)
	GetPlayers() ([]*Player, error)
	DeletePlayerById() (bool, error)
}
