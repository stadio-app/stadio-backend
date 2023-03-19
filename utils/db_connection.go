package utils

import (
	"database/sql"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/stadio-app/go-backend/ent"
	"github.com/stadio-app/go-backend/types"
)

func PostgresDnsBuilder(config types.DbConnectionOptions) string {
	ssl_mode := "enable"
	if !config.SslMode {
		ssl_mode = "disable"
	}

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
		ssl_mode,
	)
}

func PostgresDNS() string {
	return PostgresDnsBuilder(types.DbConnectionOptions{
		Host: "localhost",
		Port: "5433",
		DbName: "postgres",
		User: "postgres",
		Password: "postgres",
		SslMode: false, // TODO: disable during prod
	})
}

func DbConnection() *sql.DB {
	dns := PostgresDNS()
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
