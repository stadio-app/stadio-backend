package services

import (
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/m3-app/backend/ent"
)

type ServiceConfig struct {
	DbConn *sql.DB
	EntityManager *ent.Client
	Validator *validator.Validate
}

type Service struct {
	UserService UserService
	// list of services UserService UserService
}

func New(config ServiceConfig) Service {
	service := Service{}
	service.UserService = UserService{
		ServiceConfig: config,
	}
	return service
}
