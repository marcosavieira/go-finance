package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver  = "postgres"
	dbSourced = "postgres://gofinance_jahg_user:3CeZinb4jALqtQBbft9s4xWW8u9v6u08@dpg-clp0rgpoh6hc73bqf110-a/gofinance_jahg"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourced)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
