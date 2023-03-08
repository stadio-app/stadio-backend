package util

import (
	"database/sql"
	"fmt"
)

func DbConnection() *sql.DB {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago",
		"localhost",
		"postgres",
		"password",
		"postgres",
		"5432",
		"verify-full",
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
