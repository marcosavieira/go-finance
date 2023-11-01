package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/marcosavieira/go-finance/api"
	db "github.com/marcosavieira/go-finance/db/sqlc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading godotenv: ", err)
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbSourced := os.Getenv("DB_SOURCED")
	serverAddress := os.Getenv("SERVER_ADDRESS")

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
