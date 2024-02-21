package services

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/stadio-app/stadio-backend/types"
)

type Service struct {
	DB *sql.DB
	StructValidator *validator.Validate
	Tokens types.Tokens
}
