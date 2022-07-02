package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rhtyx/narawangsa/lib"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := lib.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = NewQueries(conn)

	os.Exit(m.Run())
}
