package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/table"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func (service Service) CreateLocation(ctx context.Context, user *gmodel.User, input gmodel.CreateLocation) (gmodel.Location, error) {
	if len(input.Instances) == 0 {
		return gmodel.Location{}, fmt.Errorf("must provide at least 1 location instance")
	}
	if len(input.Schedule) < 7 {
		return gmodel.Location{}, fmt.Errorf("minimum of 7 days of schedule is required")
	}
	if len(input.Images) == 0 || len(input.Images) > 5 {
		return gmodel.Location{}, fmt.Errorf("minimum 1 image upload, and maximum of 5 images")
	}
	
	tx, err := service.DB.BeginTx(ctx, nil)
	if err != nil {
		return gmodel.Location{}, fmt.Errorf("could not start transaction")
	}
	service.TX = tx
	defer service.TX.Rollback()

	// must create address first
	address, err := service.CreateAddress(ctx, user, *input.Address)
	if err != nil {
		return gmodel.Location{}, err
	}
	
	qb := table.Location.INSERT(
		table.Location.Name,
		table.Location.Description,
		table.Location.Type,
		table.Location.AddressID,
		table.Location.CreatedByID,
		table.Location.UpdatedByID,
	).MODEL(model.Location{
		Name: input.Name,
		Description: input.Description,
		Type: input.Type,
		AddressID: address.ID,
		CreatedByID: &user.ID,
		UpdatedByID: &user.ID,
	}).RETURNING(table.Location.AllColumns)
	
	var location gmodel.Location
	if err := qb.QueryContext(ctx, tx, &location); err != nil {
		return gmodel.Location{}, fmt.Errorf("could not create location")
	}
	location.Address = &address

	// add location instances
	location.LocationInstances, err = service.BulkCreateLocationInstances(ctx, location.ID, input.Instances)
	if err != nil {
		return gmodel.Location{}, err
	}

	// add location schedules
	location.LocationSchedule, err = service.BulkCreateLocationSchedule(ctx, location.ID, input.Schedule)
	if err != nil {
		return gmodel.Location{}, err
	}

	// upload and store images
	_, err = service.BulkUploadAndCreateLocationImages(ctx, location.ID, user, input.Images)
	if err != nil {
		tx.Rollback()
		return gmodel.Location{}, err
	}

	if err := tx.Commit(); err != nil {
		return gmodel.Location{}, fmt.Errorf("could not commit location via transaction")
	}
	return location, nil
}

func (service Service) BulkCreateLocationInstances(
	ctx context.Context, 
	location_id int64, 
	input []*gmodel.CreateLocationInstance,
) (location_instances []*gmodel.LocationInstance, err error) {
	create_location_instances := make([]model.LocationInstance, len(input))
	for i := range input {
		create_location_instances[i] = model.LocationInstance{
			LocationID: location_id,
			Name: &input[i].Name,
		}
	}
	qb := table.LocationInstance.
		INSERT(
			table.LocationInstance.LocationID,
			table.LocationInstance.Name,
		).
		MODELS(create_location_instances).
		RETURNING(table.LocationInstance.AllColumns)
	if err = qb.QueryContext(ctx, service.DbOrTxQueryable(), &location_instances); err != nil {
		return nil, err
	}
	return location_instances, nil
}

func (service Service) BulkCreateLocationSchedule(
	ctx context.Context, 
	location_id int64, 
	input []*gmodel.CreateLocationSchedule,
) ([]*gmodel.LocationSchedule, error) {
	location_schedules := make([]gmodel.LocationSchedule, len(input))
	for i := range input {
		new_location_schedule, err := service.CreateLocationSchedule(ctx, location_id, *input[i])
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

func (service Service) BulkUploadAndCreateLocationImages(
	ctx context.Context, 
	location_id int64,
	user *gmodel.User,
	image_inputs []*gmodel.CreateLocationImage,
) ([]*gmodel.LocationImage, error) {
	inserted_images := make([]*gmodel.LocationImage, len(image_inputs))
	for i, image_input := range image_inputs {
		uploaded_image, err := service.GraphImageUpload(ctx, image_input.File, uploader.UploadParams{})
		if err != nil {
			return nil, fmt.Errorf("failed to upload image %s to Cloudflare", image_input.File.Filename)
		}

		insert_qb := table.LocationImage.INSERT(
			table.LocationImage.LocationID,
			table.LocationImage.Default,
			table.LocationImage.UploadID,
			table.LocationImage.OriginalFilename,
			table.LocationImage.Caption,
			table.LocationImage.CreatedBy,
			table.LocationImage.UpdatedBy,
		).MODEL(model.LocationImage{
			LocationID: location_id,
			Default: image_input.Default,
			UploadID: uploaded_image.PublicID,
			OriginalFilename: image_input.File.Filename,
			Caption: image_input.Caption,
			CreatedBy: &user.ID,
			UpdatedBy: &user.ID,
		}).RETURNING(table.LocationImage.AllColumns)
		var image gmodel.LocationImage
		if err := insert_qb.QueryContext(ctx, service.DbOrTxQueryable(), &image); err != nil {
			return nil, err
		}
		inserted_images[i] = &image
	}
	return inserted_images, nil
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

	ls_model := model.LocationSchedule{
		LocationID: location_id,
		Day: week_day,
		Available: input.Available,
		On: input.On,
	}
	if input.From != nil && input.To != nil {
		from, err := time.Parse(time.TimeOnly, service.FromToTimeString(*input.From))
		if err != nil {
			return gmodel.LocationSchedule{}, fmt.Errorf("could not parse from date")
		}
		ls_model.From = &from
		to := service.ToDuration(*input.From, *input.To)
		ls_model.ToDuration = &to
	}

	db := service.DbOrTxQueryable()
	query_builder := table.LocationSchedule.
		INSERT(
			table.LocationSchedule.LocationID, 
			table.LocationSchedule.Day,
			table.LocationSchedule.Available,
			table.LocationSchedule.On, 
			table.LocationSchedule.From,
			table.LocationSchedule.ToDuration,
		).
		MODEL(ls_model).
		RETURNING(table.LocationSchedule.AllColumns)
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
	if !from.Before(to) {
		return false
	}

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
				postgres.OR(
					table.LocationSchedule.On.BETWEEN(postgres.Date(from.Date()), postgres.Date(to.Date())),
					table.LocationSchedule.On.IS_NULL(),
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
		// we have yet to figure out `to` is within range.
		// so we will check if `to` < `schedule.From` + `schedule.ToDuration`
		// Note: we MUST compare to with only it's time
		if schedule.From != nil && schedule.ToDuration != nil {
			schedule_to := (*schedule.From).Add(time.Hour * time.Duration(*schedule.ToDuration))
			to_time_only, _ := time.Parse(time.TimeOnly, to.Format(time.TimeOnly))
			if !to_time_only.Before(schedule_to) {
				continue
			}
		}

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

func (Service) ToDuration(from int, to int) int32 {
	if from <= to {
		return int32(to - from)
	}
	return int32((to + 24) - from)
}

func (service Service) AllLocationInstances(ctx context.Context, location_id int64) ([]model.LocationInstance, error) {
	qb := table.LocationInstance.
		SELECT(table.LocationInstance.ID).
		FROM(table.LocationInstance).
		WHERE(table.LocationInstance.LocationID.EQ(postgres.Int(location_id)))
	var location_instances []model.LocationInstance
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &location_instances); err != nil {
		return nil, err
	}
	return location_instances, nil
}

func (service Service) UnavailableLocationInstancesBetween(
	ctx context.Context,
	location_id int64,
	from time.Time,
	to time.Time,
) ([]model.LocationInstance, error) {
	// 1 minute wiggle room
	from = from.Add(time.Minute)
	to = to.Add(-1 * time.Minute)

	qb := table.LocationInstance.
		SELECT(table.LocationInstance.AllColumns).
		FROM(
			table.LocationInstance.
				INNER_JOIN(table.Location, table.Location.ID.EQ(table.LocationInstance.LocationID)).
				LEFT_JOIN(
					table.Event,
					table.Event.LocationID.EQ(table.Location.ID).
						AND(table.Event.LocationInstanceID.EQ(table.LocationInstance.ID)),
				),
		).
		WHERE(
			postgres.AND(
				table.LocationInstance.LocationID.EQ(postgres.Int(location_id)),
				postgres.OR(
					// covers cases when event from or to are within db start or end dates
					postgres.OR(
						postgres.TimestampzT(from).BETWEEN(table.Event.StartDate, table.Event.EndDate),
						postgres.TimestampzT(to).BETWEEN(table.Event.StartDate, table.Event.EndDate),
					),
					// covers cases when from or to overlap db start or end dates
					postgres.OR(
						table.Event.StartDate.BETWEEN(postgres.TimestampzT(from), postgres.TimestampzT(to)),
						table.Event.EndDate.BETWEEN(postgres.TimestampzT(from), postgres.TimestampzT(to)),
					),
				),
			),
		)
	var unavailable_instances []model.LocationInstance
	if err := qb.QueryContext(ctx, service.DbOrTxQueryable(), &unavailable_instances); err != nil {
		return nil, err
	}
	return unavailable_instances, nil
}

func (service Service) AvailableLocationInstancesBetween(ctx context.Context,
	location_id int64,
	from time.Time,
	to time.Time,
) ([]model.LocationInstance, error) {
	available_instances := []model.LocationInstance{}
	all_instances, err := service.AllLocationInstances(ctx, location_id)
	if err != nil {
		return available_instances, err
	}

	unavailable_instances, err := service.UnavailableLocationInstancesBetween(ctx, location_id, from, to)
	if err != nil {
		return available_instances, err
	}
	unavailable_instance_map := map[int64]int64{}
	for _, instance := range unavailable_instances {
		unavailable_instance_map[instance.ID] = instance.ID
	}

	for _, instance := range all_instances {
		if unavailable_instance_map[instance.ID] == 0 {
			available_instances = append(available_instances, instance)
		}
	}
	return available_instances, nil
}
