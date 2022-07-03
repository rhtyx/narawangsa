package postgres

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rhtyx/narawangsa/lib"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := lib.LoadConfig("./../../..")
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

func (q *Queries) truncate() {
	query := `
	TRUNCATE "public"."books" CASCADE;
TRUNCATE "public"."users" CASCADE;
TRUNCATE "public"."book_lists" CASCADE;
TRUNCATE "public"."user_levels" CASCADE;
TRUNCATE "public"."schema_migrations" CASCADE;
TRUNCATE "public"."categories" CASCADE;
TRUNCATE "public"."read_confirmations" CASCADE;
TRUNCATE "public"."category_books" CASCADE;
	`
	_, err := q.db.ExecContext(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
}
