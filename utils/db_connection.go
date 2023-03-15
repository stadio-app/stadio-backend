package utils

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/stadio-app/go-backend/ent"
)

func DbConnection() *sql.DB {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago",
		"localhost",
		"postgres",
		"postgres",
		"postgres",
		"5433",
		"disable",
	)
	db_conn, err := sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}

	if err := db_conn.Ping(); err != nil {
		panic(err)
	}
	return db_conn
}

func CreateEntClient(db_conn *sql.DB) *ent.Client {
	driver := entsql.OpenDB(dialect.Postgres, db_conn)
	return ent.NewClient(ent.Driver(driver))
}
