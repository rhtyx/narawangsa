package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rhtyx/narawangsa/http/server"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://narawangsa:narawangsa@localhost:5434/narawangsa_db?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	db := postgres.NewQueries(conn)
	dbTx := postgres.NewTxInContext(conn)
	server := server.New(db, dbTx)

	err = server.Start(address)
	if err != nil {
		log.Fatal(err)
	}
}
