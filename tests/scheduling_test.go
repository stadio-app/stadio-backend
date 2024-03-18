package tests

import (
	"testing"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func newint(i int) *int {
	return &i
}

func TestLocation(t *testing.T) {
	user, err := service.CreateInternalUser(ctx, gmodel.CreateAccountInput{
		Email: "location_test@thestadio.com",
		Name: "Location Test User",
		Password: "location_test_password123",
	})
	if err != nil {
		t.Fatal("could not create user")
	}

	t.Run("create location", func(t *testing.T) {
		input := gmodel.CreateLocation{
			Name: "My Soccer Field",
			Type: "Indoor Soccer Field",
			Address: &gmodel.CreateAddress{
				Latitude: 41.8413241,
				Longitude: -88.2938909,
				MapsLink: "maps.google.com",
				FullAddress: "902 Hodge Ln, Batavia, IL 60510, USA",
				CountryCode: model.CountryCodeAlpha2_Us.String(),
			},
			Schedule: []*gmodel.CreateLocationSchedule{
				{
					Day: gmodel.WeekDaySunday,
					Available: false,
				},
				{
					Day: gmodel.WeekDayMonday,
					Available: true,
					From: newint(9),
					To: newint(17),
				},
				{
					Day: gmodel.WeekDayTuesday,
					Available: true,
					From: newint(9),
					To: newint(17),
				},
				{
					Day: gmodel.WeekDayWednesday,
					Available: true,
					From: newint(9),
					To: newint(17),
				},
				{
					Day: gmodel.WeekDayThursday,
					Available: true,
					From: newint(9),
					To: newint(17),
				},
				{
					Day: gmodel.WeekDayFriday,
					Available: true,
					From: newint(9),
					To: newint(19),
				},
				{
					Day: gmodel.WeekDaySaturday,
					Available: true,
					From: newint(7),
					To: newint(17),
				},
			},
		}
		location, err := service.CreateLocation(ctx, &user, input)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(location.LocationSchedule) != len(input.Schedule) {
			t.Fatal("location.location_schedule and schedule must have the same number of values")
		}

		for i, schedule := range location.LocationSchedule {
			schedule_input := input.Schedule[i]
			if schedule_input.Available != schedule.Available {
				t.Fatal("availability does not match", schedule_input, schedule)
			}
			if schedule_input.From != nil && schedule_input.To != nil {
				if schedule.From.Hour() != *schedule_input.From {
					t.Fatal("hour does not match")
				}
				if service.ToDuration(*schedule_input.From, *schedule_input.To) != int32(*schedule.ToDuration) {
					t.Fatal("duration does not match")
				}
			}
		}
	})
}
