package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stadio-app/stadio-backend/database/jet/postgres/public/model"
	"github.com/stadio-app/stadio-backend/graph/gmodel"
)

func NewInt(i int) *int {
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
	for i := range schedule {
		schedule_ptr[i] = &schedule[i]
	}
	return schedule_ptr
}

func init_location_instances(count int) []*gmodel.CreateLocationInstance {
	instances := make([]gmodel.CreateLocationInstance, count)
	for i := range instances {
		instances[i] = gmodel.CreateLocationInstance{ Name: fmt.Sprintf("field #%d", i) }
	}
	result := make([]*gmodel.CreateLocationInstance, count)
	for i := range instances {
		result[i] = &instances[i]
	}
	return result
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
				City: "Batavia",
				AdministrativeDivision: "Illinois",
				CountryCode: model.CountryCodeAlpha2_Us.String(),
			},
			Schedule: init_schedule(9, 17, false),
			Instances: []*gmodel.CreateLocationInstance{
				{Name: "Field #1"},
			},
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

		t.Run("show all location instances", func(t *testing.T) {
			all_instances, err := service.AllLocationInstances(ctx, location.ID)
			if err != nil {
				t.Fatal("could not return all location instances", err.Error())
			}
			if len(all_instances) != len(input.Instances) {
				t.Fatal("length does not match")
			}
		})

		t.Run("show all unavailable location instances", func(t *testing.T) {
			instances, err := service.UnavailableLocationInstancesBetween(ctx, location.ID, time.Now(), time.Now().Add(time.Hour))
			if err != nil {
				t.Fatal("could not return location instances", err.Error())
			}
			if len(instances) != 0 {
				t.Fatal("there should be no other events to cause a conflict")
			}
		})

		t.Run("show all available location instances", func(t *testing.T) {
			instances, err := service.AvailableLocationInstancesBetween(ctx, location.ID, time.Now(), time.Now().Add(time.Hour))
			if err != nil {
				t.Fatal("could not return location instances", err.Error())
			}
			if len(instances) != len(input.Instances) {
				t.Fatal("there should be exactly the same instances available as the amount created")
			}
		})
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
				From: NewInt(23),
				To: NewInt(1),
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

		t.Run("override existing schedule entry", func(t *testing.T) {
			wednesday_date_only := "2024-05-01"
			custom_time, _ := time.Parse(time.DateTime, fmt.Sprintf("%s 00:00:00", wednesday_date_only))
			_, err := service.CreateLocationSchedule(ctx, location.ID, gmodel.CreateLocationSchedule{
				Day: gmodel.WeekDayWednesday,
				On: &custom_time,
				Available: false,
			})
			if err != nil {
				t.Fatal(err.Error())
			}

			_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
				Name: fmt.Sprintf("Event #7 for Location %d", location.ID),
				Type: "Pool Party",
				StartDate: custom_time.Add(time.Hour * 5),
				EndDate: custom_time.Add(time.Hour * 6),
				LocationID: location.ID,
			})
			if err == nil {
				t.Fatal("location should be unavailable for the start date")
			}

			custom_time_2, _ := time.Parse(time.DateTime, "2024-05-08 15:00:00")
			_, err = service.CreateEvent(ctx, user, gmodel.CreateEvent{
				Name: fmt.Sprintf("Event #8 for Location %d", location.ID),
				Type: "Pool Party",
				StartDate: custom_time_2,
				EndDate: custom_time_2.Add(time.Hour * 1),
				LocationID: location.ID,
			})
			if err != nil {
				t.Fatal("location should be available.", err.Error())
			}
		})
	})

	t.Run("complex scheduling", func(t *testing.T) {
		num_instances := 7
		location, err := service.CreateLocation(ctx, &user, gmodel.CreateLocation{
			Name: "My Soccer Field",
			Type: "Indoor Soccer Field",
			Address: &gmodel.CreateAddress{
				Latitude: 39.49739,
				Longitude: -109.76647,
				MapsLink: "https://www.google.com/maps/place/39%C2%B029'50.6%22N+109%C2%B045'59.3%22W/@39.49739,-109.76647,16953012m/data=!3m1!1e3!4m4!3m3!8m2!3d39.49739!4d-109.76647?entry=ttu",
				FullAddress: "Some Random Location, Vernal, UT 84078, USA",
				City: "Vernal",
				AdministrativeDivision: "Utah",
				CountryCode: model.CountryCodeAlpha2_Us.String(),
			},
			Schedule: init_schedule(7, 20, true), // 7am - 8pm
			Instances: init_location_instances(num_instances),
		})
		if err != nil {
			t.Fatal("could not create complex location")
		}
		if len(location.LocationInstances) != num_instances {
			t.Fatal("incorrect number of instances created")
		}

		now, err := time.Parse(
			time.DateTime,
			fmt.Sprintf("%s 07:00:00", time.Now().Format(time.DateOnly)),
		)
		if err != nil {
			t.Fatal(err)
		}

		t.Run("create 1 event per instance at now", func(t *testing.T) {
			for i := 0; i < num_instances; i++ {
				event, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("event #%d", i + 1),
					StartDate: now,
					EndDate: now.Add(time.Hour),
					LocationID: location.ID,
				})

				if err != nil {
					t.Fatal("could not create event", err.Error())
				}

				if event.LocationInstanceID != location.LocationInstances[i].ID {
					t.Fatal("order should be based on location_instance.id", event)
				}
			}

			t.Run("re-create 1 event per instance at now", func(t *testing.T) {
				for i := 0; i < num_instances; i++ {
					_, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
						Name: fmt.Sprintf("event #%d", i + 1),
						StartDate: now,
						EndDate: now.Add(time.Hour),
						LocationID: location.ID,
					})
	
					if err == nil {
						t.Fatal("event should already be full at the specified time")
					}
				}
			})
		})

		t.Run("create 1 event per instance 1 hour after now", func(t *testing.T) {
			cur_now := now.Add(time.Hour)
			all_instances, err := service.AllLocationInstances(ctx, location.ID)
			if err != nil {
				t.Fatal(err)
			}
			if len(all_instances) != num_instances {
				t.Fatal("incorrect number of location instances")
			}

			unavailable, err := service.UnavailableLocationInstancesBetween(ctx, location.ID, cur_now, cur_now.Add(time.Hour))
			if err != nil {
				t.Fatal(err)
			}
			if len(unavailable) != 0 {
				t.Fatal("there should be no unavailable location instances for this date range")
			}

			for i := 0; i < num_instances; i++ {
				event, err := service.CreateEvent(ctx, user, gmodel.CreateEvent{
					Name: fmt.Sprintf("event #%d", i + 1),
					StartDate: cur_now,
					EndDate: cur_now.Add(time.Hour),
					LocationID: location.ID,
				})

				if err != nil {
					t.Fatal("could not create event", err.Error())
				}

				if event.LocationInstanceID != location.LocationInstances[i].ID {
					t.Fatal("order should be based on location_instance.id", event)
				}
			}

			unavailable, err = service.UnavailableLocationInstancesBetween(ctx, location.ID, cur_now, cur_now.Add(time.Hour))
			if err != nil {
				t.Fatal(err)
			}
			if len(unavailable) != num_instances {
				t.Fatalf("there should be %d unavailable instances for the date range", num_instances)
			}

			available, err := service.AvailableLocationInstancesBetween(ctx, location.ID, cur_now, cur_now.Add(time.Hour))
			if err != nil {
				t.Fatal(err)
			}
			if len(available) != 0 {
				t.Fatal("there should be no available location instances")
			}
		})
	})
}
