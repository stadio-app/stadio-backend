package services

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	DB *sql.DB
	StructValidator *validator.Validate
}
