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
	dbSourced = "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable"
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
