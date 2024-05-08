package services

import (
	"context"
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func (service Service) CreateEvent(ctx context.Context, user gmodel.User, input gmodel.CreateEvent) (gmodel.EventShallow, error) {
	db := service.DbOrTxQueryable()
	if !service.LocationIdExists(ctx, input.LocationID) {
		return gmodel.EventShallow{}, fmt.Errorf("location id is invalid")
	}
	if !service.LocationScheduleAvailableBetween(ctx, input.LocationID, input.StartDate, input.EndDate) {
		return gmodel.EventShallow{}, fmt.Errorf("location is unavailable at the selected date range")
	}

	location_instances, err := service.AvailableLocationInstancesBetween(ctx, input.LocationID, input.StartDate, input.EndDate)
	if err != nil {
		return gmodel.EventShallow{}, err
	}
	if len(location_instances) == 0 {
		return gmodel.EventShallow{}, fmt.Errorf("location is full at this time")
	}

	qb := table.Event.
		INSERT(
			table.Event.LocationID,
			table.Event.Name,
			table.Event.Description,
			table.Event.Type,
			table.Event.StartDate,
			table.Event.EndDate,
			table.Event.LocationInstanceID,
			table.Event.CreatedByID,
			table.Event.UpdatedByID,
		).
		MODEL(model.Event{
			LocationID:         &input.LocationID,
			Name:               input.Name,
			Description:        input.Description,
			Type:               &input.Type,
			StartDate:          input.StartDate,
			EndDate:            input.EndDate,
			LocationInstanceID: &location_instances[0].ID,
			CreatedByID:        &user.ID,
			UpdatedByID:        &user.ID,
		}).RETURNING(table.Event.AllColumns)
	var new_event gmodel.EventShallow
	if err := qb.QueryContext(ctx, db, &new_event); err != nil {
		return gmodel.EventShallow{}, err
	}
	return new_event, nil
}

func (service Service) FindAllEvents(ctx context.Context, filter gmodel.AllEventsFilter) ([]gmodel.Event, error) {
	created_by_user_table := table.User.AS("created_by_user")
	updated_by_user_table := table.User.AS("updated_by_user")
	coordinates_col_name := fmt.Sprintf("%s.%s", table.Address.Coordinates.TableName(), table.Address.Coordinates.Name())
	// dynamic column used to sort results by the calculated distance
	distance_col_name := fmt.Sprintf("%s.%s", table.Address.Coordinates.TableName(), "distance")

	qb := table.Event.
		SELECT(
			table.Event.AllColumns,
			table.Location.AllColumns,
			table.Address.AllColumns,
			table.Country.Name,
			table.LocationImage.AllColumns,
			created_by_user_table.ID,
			created_by_user_table.Name,
			created_by_user_table.Avatar,
			updated_by_user_table.ID,
			updated_by_user_table.Name,
			updated_by_user_table.Avatar,
			postgres.RawString(
				fmt.Sprintf(
					"ST_Distance(%s, 'SRID=4326;POINT(%f %f)'::geometry)",
					coordinates_col_name,
					filter.Longitude,
					filter.Latitude,
				),
			).AS(distance_col_name),
		).
		FROM(
			table.Event.
				INNER_JOIN(table.Location, table.Location.ID.EQ(table.Event.LocationID)).
				INNER_JOIN(table.Address, table.Address.ID.EQ(table.Location.AddressID)).
				INNER_JOIN(table.Country, table.Country.Code.EQ(table.Address.CountryCode)).
				LEFT_JOIN(table.LocationImage, table.LocationImage.LocationID.EQ(table.Location.ID)).
				LEFT_JOIN(created_by_user_table, created_by_user_table.ID.EQ(table.Event.CreatedByID)).
				LEFT_JOIN(updated_by_user_table, updated_by_user_table.ID.EQ(table.Event.CreatedByID)),
		).
		WHERE(
			postgres.AND(
				table.Address.CountryCode.EQ(postgres.NewEnumValue(filter.CountryCode)),
				table.Event.StartDate.GT_EQ(postgres.TimestampzT(filter.StartDate)),
				table.Event.EndDate.LT_EQ(postgres.TimestampzT(filter.EndDate)),
				postgres.RawBool(
					fmt.Sprintf(
						"ST_DWithin(%s, 'POINT(%f %f)'::geometry, %d, TRUE)",
						coordinates_col_name,
						filter.Longitude,
						filter.Latitude,
						filter.RadiusMeters,
					),
				),
			),
		).
		ORDER_BY(
			postgres.FloatColumn(distance_col_name).ASC(),
			table.Event.StartDate.ASC(),
			table.Event.CreatedAt.ASC(),
		)
	db := service.DbOrTxQueryable()
	var events []gmodel.Event
	if err := qb.QueryContext(ctx, db, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (service Service) MyEvents(ctx context.Context, user_id int64) ([]gmodel.Event, error) {
	query := table.Event.
		SELECT(
			table.Event.AllColumns,
			table.Participant.AllColumns,
			table.Location.AllColumns,
			table.Address.AllColumns,
			table.Country.Name,
			table.LocationImage.AllColumns,
			table.User.ID.AS("created_by_user_id"),
			table.User.Name.AS("created_by_name"),
			table.User.Avatar.AS("created_by_avatar"),
		).FROM(
		table.Event.
			INNER_JOIN(table.Location, table.Location.ID.EQ(table.Event.LocationID)).
			INNER_JOIN(table.Address, table.Address.ID.EQ(table.Location.AddressID)).
			INNER_JOIN(table.Country, table.Country.Code.EQ(table.Address.CountryCode)).
			LEFT_JOIN(table.Participant, table.Participant.EventID.EQ(table.Event.ID)).
			LEFT_JOIN(table.LocationImage, table.LocationImage.LocationID.EQ(table.Location.ID)).
			LEFT_JOIN(table.User, table.User.ID.EQ(table.Event.CreatedByID)),
	).WHERE(
		table.Participant.UserID.EQ(postgres.Int64(user_id)),
	).ORDER_BY(
		table.Event.StartDate.ASC(),
	)

	db := service.DbOrTxQueryable()
	var events []gmodel.Event
	if err := query.QueryContext(ctx, db, &events); err != nil {
		return nil, err
	}
	return events, nil
}
