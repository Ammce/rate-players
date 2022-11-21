package postgres

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

var dbConnection *sqlx.DB

func Connect() {
	var connStr = "postgres://postgres:postgres@docker.for.mac.localhost:5432/rate_players?sslmode=disable"
	cn, err := sqlx.Connect("postgres", connStr)
	dbConnection = cn
	if err != nil {
		fmt.Println("Error while connecting")
		log.Fatalln(err)
	}
}

func GetDBInstance() *sqlx.DB {
	return dbConnection
}
