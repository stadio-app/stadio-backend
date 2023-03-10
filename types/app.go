package types

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/m3-app/backend/ent"
)

type AppContext struct {
	DbConn *sql.DB
	Tokens *Tokens
	Server *fiber.App
	Validator *validator.Validate
	EntityManager *ent.Client
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