package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/marcosavieira/go-finance/api"
	db "github.com/marcosavieira/go-finance/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSourced     = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSourced)
	if err != nil {
		log.Fatal("cannot connect to database", err)
		return
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start api", err)
		return
	}
}
