package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	var connStr = "postgres://postgres:postgres@docker.for.mac.localhost:5432/rate_players?sslmode=disable"
	connection, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println("Error while connecting")
		log.Fatalln(err)
	}
	return connection
}
