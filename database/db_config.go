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

func DbConfig() (DbConnection, error) {
	var db_config DbConnection
	err := utils.FileMapper("db_config.json", &db_config)
	if os.Getenv("PGENV") == "production" {
		return ProdDbConfig(), nil
	}
	return db_config, err
}

func ProdDbConfig() DbConnection {
	port, err := strconv.Atoi(os.Getenv("PGPORT"))
	if err != nil {
		panic("Error establishing DB connection, no port found.")
	}
	return DbConnection{
		Host:     os.Getenv("PGHOST"),
		DbName:   os.Getenv("PGDATABASE"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("PGPASSWORD"),
		Port:     uint16(port),
	}
}

func CreateDbConnection(options DbConnection) (*sql.DB, error) {
	return sql.Open("postgres", DbConnectionString(options))
}

func NewDbConnection() (*sql.DB, error) {
	config, err := DbConfig()
	if err != nil {
		return nil, err
	}
	return CreateDbConnection(config)
}
