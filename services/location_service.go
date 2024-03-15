package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-jet/jet/v2/postgres"
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
		table.Location.CreatedByID,
		table.Location.UpdatedByID,
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
		table.LocationSchedule.ToDuration,
	).VALUES(
		location_id,
		week_day,
		input.On,
		service.FromToTimeString(*input.From),
		service.ToDuration(*input.From, *input.To),
	).RETURNING(table.LocationSchedule.AllColumns)
	var location_schedule gmodel.LocationSchedule
	if err := query_builder.QueryContext(ctx, db, &location_schedule); err != nil {
		return gmodel.LocationSchedule{}, err
	}
	return location_schedule, nil
}

func (service Service) LocationIdExists(ctx context.Context, location_id int64) bool {
	qb := table.Location.
		SELECT(table.Location.ID).
		FROM(table.Location).
		WHERE(table.Location.ID.EQ(postgres.Int64(location_id))).
		LIMIT(1)
	var dest struct{ ID string }
	return qb.QueryContext(ctx, service.DbOrTxQueryable(), &dest) == nil
}

func (service Service) FindLocationById(ctx context.Context, location_id int64) (gmodel.Location, error) {
	qb := table.Location.
		SELECT(
			table.Location.AllColumns, 
			table.LocationSchedule.AllColumns,
		).
		FROM(
			table.Location.
			INNER_JOIN(
				table.LocationSchedule, 
				table.LocationSchedule.LocationID.EQ(table.Location.ID),
			),
		).
		WHERE(table.Location.ID.EQ(postgres.Int64(location_id)))
	db := service.DbOrTxQueryable()
	var location gmodel.Location
	if err := qb.QueryContext(ctx, db, &location); err != nil {
		return gmodel.Location{}, err
	}
	return location, nil
}

// Queries the DB to check if any entries exists in `location_schedule`
// where the given date range is within the schedule date range.
func (service Service) LocationScheduleAvailableBetween(ctx context.Context, location_id int64, from time.Time, to time.Time) bool {
	from_dow := strings.ToUpper(from.Weekday().String())
	to_duration := to.Sub(from).Hours()

	qb := table.LocationSchedule.
		SELECT(table.LocationSchedule.AllColumns).
		FROM(table.LocationSchedule).
		WHERE(
			postgres.AND(
				table.LocationSchedule.LocationID.EQ(postgres.Int(location_id)),
				table.LocationSchedule.Day.EQ(postgres.NewEnumValue(from_dow)),
				postgres.AND(
					postgres.TimeT(from).GT_EQ(table.LocationSchedule.From),
					postgres.Int32(int32(to_duration)).LT_EQ(table.LocationSchedule.ToDuration),
				),
			),
		).
		ORDER_BY(table.LocationSchedule.CreatedAt.DESC())
	var available_schedules []gmodel.LocationSchedule
	db := service.DbOrTxQueryable()
	if err := qb.QueryContext(ctx, db, &available_schedules); err != nil {
		return false
	}

	// check if selected schedules are available
	for _, schedule := range available_schedules {
		if schedule.On == nil {
			return schedule.Available
		}
		// schedule.On is defined so return it's availability if date matches
		if schedule.On.Format(time.DateOnly) == from.Format(time.DateOnly) {
			return schedule.Available
		}
	}
	return false
}

func (Service) FromToTimeString(from int) string {
	return fmt.Sprintf("%d:00:00", from)
}

func (Service) ToDuration(from int, to int) int {
	if from <= to {
		return to - from
	}
	return (to + 24) - from
}
