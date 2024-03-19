package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func newint(i int) *int {
	return &i
}

func init_schedule(from int, to int, open_on_weekends bool) []*gmodel.CreateLocationSchedule {
	schedule := make([]gmodel.CreateLocationSchedule, 7)
	for i, week_day := range gmodel.AllWeekDay {
		if !open_on_weekends && (week_day == gmodel.WeekDaySaturday || week_day == gmodel.WeekDaySunday) {
			schedule[i] = gmodel.CreateLocationSchedule{
				Day: week_day,
				Available: false,
			}
			continue
		}
		schedule[i] = gmodel.CreateLocationSchedule{
			Day: week_day,
			Available: true,
			From: &from,
			To: &to,
		}
	}
	schedule_ptr := make([]*gmodel.CreateLocationSchedule, 7)
	for i, _ := range schedule {
		schedule_ptr[i] = &schedule[i]
	}
	return schedule_ptr
}

func TestLocation(t *testing.T) {
	user, err := service.CreateInternalUser(ctx, gmodel.CreateAccountInput{
		Email: "location_test@thestadio.com",
		Name: "Location Test User",
		Password: "location_test_password123",
	})
	if err != nil {
		t.Fatal("could not create user.", err.Error())
	}
	var location gmodel.Location

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
			Schedule: init_schedule(9, 17, false),
		}
		if location, err = service.CreateLocation(ctx, &user, input); err != nil {
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

	t.Run("create event", func(t *testing.T) {
		var event1 gmodel.EventShallow
		monday_date_only := "2024-03-18"

		t.Run("within normal business hours", func(t *testing.T) {
			var err error
			start_date, err := time.Parse(time.DateTime, fmt.Sprintf("%s 12:30:00", monday_date_only))
			if err != nil {
				t.Fatal("could not convert time.", err.Error())
			}
			event1, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
				Name: fmt.Sprintf("Event #1 for Location %d", location.ID),
				Type: "Pool Party",
				StartDate: start_date,
				EndDate: start_date.Add(time.Hour),
				LocationID: location.ID,
			})
			if err != nil {
				t.Fatal("could not create event.", err.Error())
			}
			if event1.Approved {
				t.Fatal("event should not be approved by default")
			}
		})

		t.Run("weekends while closed", func(t *testing.T) {
			start_date, _ := time.Parse(time.DateTime, "2024-03-16 10:00:00")
			_, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
				Name: fmt.Sprintf("Saturday event for Location %d", location.ID),
				Type: "Pool Party",
				StartDate: start_date,
				EndDate: start_date.Add(time.Hour),
				LocationID: location.ID,
			})
			if err == nil {
				t.Fatal("location is unavailable during Saturdays")
			}

			start_date, _ = time.Parse(time.DateTime, "2024-03-17 14:00:00")
			_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
				Name: fmt.Sprintf("Sunday event for Location %d", location.ID),
				Type: "Pool Party",
				StartDate: start_date,
				EndDate: start_date.Add(time.Hour),
				LocationID: location.ID,
			})
			if err == nil {
				t.Fatal("location is unavailable during Sundays")
			}
		})

		t.Run("same day events", func(t *testing.T) {
			t.Run("unique time slot within business hours", func(t *testing.T) {
				start_date, _ := time.Parse(time.DateTime, fmt.Sprintf("%s 10:00:00", monday_date_only))
				_, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #2 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: start_date,
					EndDate: start_date.Add(time.Hour),
					LocationID: location.ID,
				})
				if err != nil {
					t.Fatal("could not create event.", err.Error())
				}
			})

			t.Run("unique time slot outside business hours", func(t *testing.T) {
				start_date, _ := time.Parse(time.DateTime, fmt.Sprintf("%s 22:00:00", monday_date_only))
				_, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #3 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: start_date,
					EndDate: start_date.Add(time.Hour),
					LocationID: location.ID,
				})
				if err == nil {
					t.Fatal("event should not be created. time slot falls outside of business hours")
				}
			})

			t.Run("conflicting time slot", func(t *testing.T) {
				start_date, _ := time.Parse(time.DateTime, fmt.Sprintf("%s 10:00:00", monday_date_only))
				_, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #4 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: start_date,
					EndDate: start_date.Add(time.Hour),
					LocationID: location.ID,
				})
				if err == nil {
					t.Fatal("should not create event. this event conflicts with event #2")
				}
			})
		})

		t.Run("special schedule entry", func(t *testing.T) {
			friday_date_only := "2024-03-22"
			custom_time, _ := time.Parse(time.DateTime, fmt.Sprintf("%s 23:00:00", friday_date_only))
			_, err := service.CreateLocationSchedule(ctx, location.ID, gmodel.CreateLocationSchedule{
				Day: gmodel.WeekDayFriday,
				On: &custom_time,
				From: newint(23),
				To: newint(1),
				Available: true,
			})
			if err != nil {
				t.Fatal("could not create custom schedule.", err.Error())
			}

			t.Run("within custom schedule range", func(t *testing.T) {
				_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #5 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: custom_time,
					EndDate: custom_time.Add(time.Hour),
					LocationID: location.ID,
				})
				if err != nil {
					t.Fatal("could not add event to custom schedule.", err.Error())
				}
			})

			t.Run("overlapping custom schedule range", func(t *testing.T) {
				_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #6 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: custom_time.Add(time.Hour),
					EndDate: custom_time.Add(time.Hour * 3),
					LocationID: location.ID,
				})
				if err == nil {
					t.Fatal("this event's end date overlaps the schedule")
				}
			})

			t.Run("outside custom schedule range", func(t *testing.T) {
				_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("Event #6 for Location %d", location.ID),
					Type: "Pool Party",
					StartDate: custom_time.Add(time.Hour * 5),
					EndDate: custom_time.Add(time.Hour * 10),
					LocationID: location.ID,
				})
				if err == nil {
					t.Fatal("event is outside the schedule")
				}
			})
		})
	})
}
