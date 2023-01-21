package import_data

type ImportPlayerService interface {
	ImportPlayersForTeam(teamId int8) bool
}
