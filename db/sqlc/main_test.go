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
	dbSourced = "postgres://root:XWCqqbjbDUoLsgUNdSSfNU4ZlZCmvS4K@dpg-clp1251oh6hc73bqk3d0-a.oregon-postgres.render.com/root_li8f"
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
