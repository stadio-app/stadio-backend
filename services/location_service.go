package services

import (
	"context"
	"fmt"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func (service Service) CreateLocation(ctx context.Context, user *gmodel.User, input gmodel.CreateLocation) (gmodel.Location, error) {
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return gmodel.Location{}, fmt.Errorf("could not start transaction")
	}
	service.TX = tx

	// must create address first
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

	// add location schedules
	location.LocationSchedule, err = service.BulkCreateLocationSchedule(ctx, location.ID, input.Schedule)
	if err != nil {
		tx.Rollback()
		return gmodel.Location{}, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return gmodel.Location{}, fmt.Errorf("could not commit location via transaction")
	}
	return location, nil
}

func (service Service) BulkCreateLocationSchedule(
	ctx context.Context, 
	location_id int64, 
	input []*gmodel.CreateLocationSchedule,
) ([]*gmodel.LocationSchedule, error) {
	location_schedules := make([]gmodel.LocationSchedule, len(input))
	for i, schedule_input := range input {
		new_location_schedule, err := service.CreateLocationSchedule(ctx, location_id, *schedule_input)
		if err != nil {
			return nil, fmt.Errorf("could not create location schedule. %s", err.Error())
		}
		location_schedules[i] = new_location_schedule
	}
	location_schedule_ptrs := make([]*gmodel.LocationSchedule, len(location_schedules))
	for i := range location_schedules {
		location_schedule_ptrs[i] = &location_schedules[i]
	}
	return location_schedule_ptrs, nil
}

func (service Service) CreateLocationSchedule(
	ctx context.Context, 
	location_id int64, 
	input gmodel.CreateLocationSchedule,
) (gmodel.LocationSchedule, error) {
	var week_day model.WeekDay
	if week_day.Scan(input.Day.String()) != nil {
		return gmodel.LocationSchedule{}, fmt.Errorf("parsing error with week day enum")
	}
	if input.From != nil && input.To == nil {
		return gmodel.LocationSchedule{}, fmt.Errorf("schedule field 'to' is not defined")
	}

	db := service.DbOrTxQueryable()
	query_builder := table.LocationSchedule.INSERT(
		table.LocationSchedule.LocationID,
		table.LocationSchedule.Day,
		table.LocationSchedule.On,
		table.LocationSchedule.From,
		table.LocationSchedule.To,
	).VALUES(
		location_id,
		week_day,
		input.On,
		input.From,
		input.To,
	).RETURNING(table.LocationSchedule.AllColumns)
	var location_schedule gmodel.LocationSchedule
	if err := query_builder.QueryContext(ctx, db, &location_schedule); err != nil {
		return gmodel.LocationSchedule{}, err
	}
	return location_schedule, nil
}
