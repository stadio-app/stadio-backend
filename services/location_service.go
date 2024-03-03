package services

import (
	"context"
	"fmt"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func (service Service) CreateLocation(ctx context.Context, user *gmodel.User, input gmodel.CreateLocation) (gmodel.Location, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return gmodel.Location{}, fmt.Errorf("could not start transaction")
	}
	service.TX = tx
	address, err := service.CreateAddress(ctx, user, *input.Address)
	if err != nil {
		tx.Rollback()
		return gmodel.Location{}, err
	}
	
	qb := table.Location.INSERT(
		table.Location.Name,
		table.Location.Description,
		table.Location.Type,
		table.Location.AddressID,
		table.Location.CreatedBy,
		table.Location.UpdatedBy,
	).VALUES(
		input.Name,
		input.Description,
		input.Type,
		address.ID,
		user.ID,
		user.ID,
	).RETURNING(table.Location.AllColumns)
	
	var location gmodel.Location
	if err := qb.QueryContext(ctx, tx, &location); err != nil {
		tx.Rollback()
		return gmodel.Location{}, fmt.Errorf("could not create location")
	}
	location.Address = &address
	location.CreatedBy = user
	location.UpdatedBy = user
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return gmodel.Location{}, fmt.Errorf("could not commit location via transaction")
	}
	return location, nil
}
