package db

import (
	"database/sql"
	_ "github.com/lib/pq" // postgres driver, imported but not referenced in code, hence the leading _
	"log"
	"testing"
)

const (
	dbDriver   = "postgres"
	dataSource = "postgres://root:123@localhost:5432/simple_bank?sslmode=disable"
)

var testQuery *Queries

// the main entry point for all tests in package db
func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dataSource)
	if err != nil {
		log.Fatal("Cannot establish database connection.")
	}
	testQuery = New(conn)

	m.Run()
}
