package types

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/m3-app/backend/ent"
)

type AppContext struct {
	DbConn *sql.DB
	Tokens *Tokens
	Validator *validator.Validate
	EntityManager *ent.Client
	Port string
}

type DbConnectionOptions struct {
	Host string
	User string
	Password string
	DbName string
	Port string
	SslMode bool
	DisableLogger bool
}