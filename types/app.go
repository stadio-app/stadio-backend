package types

import (
	"database/sql"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-playground/validator"
	"github.com/stadio-app/stadio-backend/ent"
)

type AuthKeyType string
const AuthKey AuthKeyType = "auth"
const AuthHeader string = "Authorization"

type AppContext struct {
	DbConn *sql.DB
	Tokens *Tokens
	Validator *validator.Validate
	EntityDriver *entsql.Driver
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
}

type FullName struct {
	FirstName string
	MiddleName *string
	LastName string
}
