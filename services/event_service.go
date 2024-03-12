package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func (service Service) CreateEvent(ctx context.Context, user gmodel.User, input gmodel.CreateEvent) (gmodel.Event, error) {
	db := service.DbOrTxQueryable()
	location, err := service.FindLocationById(ctx, input.LocationID)
	if err != nil {
		return gmodel.Event{}, fmt.Errorf("location id is invalid")
	}

	if !service.LocationScheduleAvailableBetween(ctx, location.ID, input.StartDate, input.EndDate) {
		return gmodel.Event{}, fmt.Errorf("location is unavailable at the selected date range")
	}

	if service.EventTimingCollides(ctx, location.ID, input.StartDate, input.EndDate) {
		return gmodel.Event{}, fmt.Errorf("time slot already taken by another event")
	}

	qb := table.Event.
		INSERT(
			table.Event.LocationID,
			table.Event.Name,
			table.Event.Description,
			table.Event.Type,
			table.Event.StartDate,
			table.Event.EndDate,
			table.Event.CreatedByID,
			table.Event.UpdatedByID,
		).
		VALUES(
			location.ID,
			input.Name,
			input.Description,
			input.Type,
			input.StartDate,
			input.EndDate,
			user.ID,
			user.ID,
		).RETURNING(table.Event.AllColumns)
	var new_event gmodel.Event
	if err := qb.QueryContext(ctx, db, &new_event); err != nil {
		return gmodel.Event{}, err
	}
	new_event.CreatedBy = &user
	new_event.UpdatedBy = &user
	new_event.Location = &location
	return new_event, nil
}

func (service Service) EventTimingCollides(ctx context.Context, location_id int64, from time.Time, to time.Time) bool {
	qb := table.Event.
		SELECT(table.Event.ID).
		FROM(table.Event).
		WHERE(
			postgres.AND(
				table.Event.LocationID.EQ(postgres.Int(location_id)),
				postgres.OR(
					postgres.OR(
						postgres.TimestampzT(from).BETWEEN(table.Event.StartDate, table.Event.EndDate),
						postgres.TimestampzT(to).BETWEEN(table.Event.StartDate, table.Event.EndDate),
					),
					postgres.OR(
						table.Event.StartDate.BETWEEN(postgres.TimestampzT(from), postgres.TimestampzT(to)),
						table.Event.EndDate.BETWEEN(postgres.TimestampzT(from), postgres.TimestampzT(to)),
					),
				),
			),
		)
	var conflicting_events []int
	db := service.DbOrTxQueryable()
	if err := qb.QueryContext(ctx, db, &conflicting_events); err != nil {
		log.Println(err)
		return true
	}
	log.Println(conflicting_events)
	return len(conflicting_events) != 0
}
