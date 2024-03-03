package services

import (
	"database/sql"

	"github.com/go-jet/jet/v2/qrm"
	"github.com/go-playground/validator/v10"
	"github.com/stadio-app/stadio-backend/types"
)

type Service struct {
	DB *sql.DB
	TX *sql.Tx
	StructValidator *validator.Validate
	Tokens *types.Tokens
}

// Returns a transaction if present.
// Otherwise returns the Database connection instance as qrm.Queryable
func (service *Service) DbOrTxQueryable() qrm.Queryable {
	if service.TX != nil {
		return service.TX
	}
	return service.DB
}
