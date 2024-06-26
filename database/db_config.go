package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/stadio-app/stadio-backend/utils"
)

type DbConnection struct {
	DbName   string `json:"dbName"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	SslMode  bool   `json:"sslMode"`
}

func DbConnectionString(options DbConnection) string {
	sslmode_val := "enable"
	if !options.SslMode {
		sslmode_val = "disable"
	}
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Chicago",
		options.Host,
		options.Username,
		options.Password,
		options.DbName,
		strconv.Itoa(int(options.Port)),
		sslmode_val,
	)
	return dns
}

// This method should only be used for local development
func DbConfig() (db_config DbConnection, err error) {
	err = utils.FileMapper("db_config.json", &db_config)
	return db_config, err
}

func CreateDbConnection(options DbConnection) (*sql.DB, error) {
	if db_url, exists := os.LookupEnv("PG_DATABASE_URL"); exists {
		return sql.Open("postgres", db_url)
	}
	return sql.Open("postgres", DbConnectionString(options))
}

func NewDbConnection() (*sql.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	return CreateDbConnection(config)
}
