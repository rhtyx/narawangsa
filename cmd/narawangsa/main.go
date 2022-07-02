package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rhtyx/narawangsa/http/server"
	"github.com/rhtyx/narawangsa/internal/storage/postgres"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

func main() {
	config, err := lib.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		log.Fatal("cannot create tokenMaker: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	db := postgres.NewQueries(conn)
	dbTx := postgres.NewTxInContext(conn)
	server := server.New(db, dbTx, config, tokenMaker)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
