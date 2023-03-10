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

type ServiceBase struct {
	ServiceConfig
	Table string
}

type Service struct {
	ServiceConfig
	UserService UserService
	// list of services UserService UserService
}

func CreateService(config ServiceConfig, table string) ServiceBase {
	return ServiceBase{
		ServiceConfig: config,
		Table: table,
	}
}

func New(config ServiceConfig) Service {
	service := Service{
		ServiceConfig: config,
	}
	service.UserService = UserService{
		ServiceBase: CreateService(config, "users"),
	}
	return service
}
