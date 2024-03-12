// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gmodel

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Address struct {
	ID          int64        `json:"id" sql:"primary_key"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Latitude    float64      `json:"latitude"`
	Longitude   float64      `json:"longitude"`
	MapsLink    string       `json:"mapsLink"`
	FullAddress string       `json:"fullAddress"`
	CountryCode string       `json:"countryCode"`
	Country     *string      `json:"country,omitempty" alias:"country.name"`
	CreatedByID *int64       `json:"createdById,omitempty"`
	CreatedBy   *UserShallow `json:"createdBy,omitempty"`
	UpdatedByID *int64       `json:"updatedById,omitempty"`
	UpdatedBy   *UserShallow `json:"updatedBy,omitempty"`
}

type Auth struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type CreateAccountInput struct {
	Email       string  `json:"email" validate:"required,email"`
	PhoneNumber *string `json:"phoneNumber,omitempty" validate:"omitempty,e164"`
	Name        string  `json:"name" validate:"required"`
	Password    string  `json:"password" validate:"required"`
}

type CreateAddress struct {
	Latitude    float64 `json:"latitude" validate:"required,latitude"`
	Longitude   float64 `json:"longitude" validate:"required,longitude"`
	MapsLink    string  `json:"mapsLink" validate:"required,http_url"`
	FullAddress string  `json:"fullAddress" validate:"required,contains"`
	CountryCode string  `json:"countryCode" validate:"iso3166_1_alpha2"`
}

type CreateEvent struct {
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Type        string    `json:"type"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	LocationID  int64     `json:"locationId"`
}

type CreateLocation struct {
	Name        string                    `json:"name" validate:"required"`
	Description *string                   `json:"description,omitempty"`
	Type        string                    `json:"type" validate:"required"`
	Address     *CreateAddress            `json:"address" validate:"required"`
	Schedule    []*CreateLocationSchedule `json:"schedule"`
}

type CreateLocationSchedule struct {
	Day       WeekDay    `json:"day" validate:"required"`
	On        *time.Time `json:"on,omitempty"`
	From      *int       `json:"from,omitempty" validate:"gte=0,lt=24"`
	To        *int       `json:"to,omitempty" validate:"gte=0,lt=24"`
	Available bool       `json:"available" validate:"required"`
}

type Event struct {
	ID          int64     `json:"id" sql:"primary_key"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Type        string    `json:"type"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	LocationID  int64     `json:"locationId"`
	Location    *Location `json:"location,omitempty"`
	CreatedByID *int64    `json:"createdById,omitempty"`
	CreatedBy   *User     `json:"createdBy,omitempty"`
	UpdatedByID *int64    `json:"updatedById,omitempty"`
	UpdatedBy   *User     `json:"updatedBy,omitempty"`
	Approved    bool      `json:"approved"`
}

type EventShallow struct {
	ID          int64     `json:"id" sql:"primary_key" alias:"event.id"`
	CreatedAt   time.Time `json:"createdAt" alias:"event.created_at"`
	UpdatedAt   time.Time `json:"updatedAt" alias:"event.updated_at"`
	Name        string    `json:"name" alias:"event.name"`
	Description *string   `json:"description,omitempty" alias:"event.description"`
	Type        string    `json:"type" alias:"event.type"`
	StartDate   time.Time `json:"startDate" alias:"event.start_date"`
	EndDate     time.Time `json:"endDate" alias:"event.end_date"`
	LocationID  int64     `json:"locationId" alias:"event.location_id"`
	CreatedByID *int64    `json:"createdById,omitempty" alias:"event.created_by_id"`
	UpdatedByID *int64    `json:"updatedById,omitempty" alias:"event.updated_by_id"`
	Approved    bool      `json:"approved" alias:"event.approved"`
}

type Location struct {
	ID               int64               `json:"id" sql:"primary_key"`
	CreatedAt        time.Time           `json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
	Name             string              `json:"name"`
	Description      *string             `json:"description,omitempty"`
	Type             string              `json:"type"`
	OwnerID          *int64              `json:"ownerId,omitempty"`
	Owner            *Owner              `json:"owner,omitempty"`
	AddressID        int64               `json:"addressId"`
	Address          *Address            `json:"address,omitempty"`
	Deleted          bool                `json:"deleted"`
	Status           string              `json:"status"`
	CreatedByID      *int64              `json:"createdById,omitempty"`
	CreatedBy        *User               `json:"createdBy,omitempty"`
	UpdatedByID      *int64              `json:"updatedById,omitempty"`
	UpdatedBy        *User               `json:"updatedBy,omitempty"`
	LocationSchedule []*LocationSchedule `json:"locationSchedule"`
}

type LocationSchedule struct {
	ID         int64      `json:"id" sql:"primary_key"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	LocationID int64      `json:"locationId"`
	Location   *Location  `json:"location,omitempty"`
	Day        WeekDay    `json:"day"`
	On         *time.Time `json:"on,omitempty"`
	From       *int       `json:"from,omitempty"`
	To         *int       `json:"to,omitempty"`
	Available  bool       `json:"available"`
}

type Mutation struct {
}

type Owner struct {
	ID         int64        `json:"id" sql:"primary_key"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	FirstName  string       `json:"firstName"`
	MiddleName *string      `json:"middleName,omitempty"`
	LastName   string       `json:"lastName"`
	FullName   string       `json:"fullName"`
	Verified   bool         `json:"verified"`
	UserID     int64        `json:"userId"`
	User       *UserShallow `json:"user,omitempty"`
}

type Query struct {
}

type User struct {
	ID           int64             `json:"id" sql:"primary_key"`
	CreatedAt    time.Time         `json:"createdAt"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	Email        string            `json:"email"`
	PhoneNumber  *string           `json:"phoneNumber,omitempty"`
	Name         string            `json:"name"`
	Avatar       *string           `json:"avatar,omitempty"`
	BirthDate    *time.Time        `json:"birthDate,omitempty"`
	Bio          *string           `json:"bio,omitempty"`
	Active       *bool             `json:"active,omitempty"`
	AuthPlatform *AuthPlatformType `json:"authPlatform,omitempty"`
	AuthStateID  int64             `json:"authStateId" alias:"auth_state.id"`
}

type UserShallow struct {
	ID        int64     `json:"id" sql:"primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Avatar    *string   `json:"avatar,omitempty"`
	Bio       *string   `json:"bio,omitempty"`
	Active    *bool     `json:"active,omitempty"`
}

type AuthPlatformType string

const (
	AuthPlatformTypeInternal AuthPlatformType = "INTERNAL"
	AuthPlatformTypeApple    AuthPlatformType = "APPLE"
	AuthPlatformTypeGoogle   AuthPlatformType = "GOOGLE"
)

var AllAuthPlatformType = []AuthPlatformType{
	AuthPlatformTypeInternal,
	AuthPlatformTypeApple,
	AuthPlatformTypeGoogle,
}

func (e AuthPlatformType) IsValid() bool {
	switch e {
	case AuthPlatformTypeInternal, AuthPlatformTypeApple, AuthPlatformTypeGoogle:
		return true
	}
	return false
}

func (e AuthPlatformType) String() string {
	return string(e)
}

func (e *AuthPlatformType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuthPlatformType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuthPlatformType", str)
	}
	return nil
}

func (e AuthPlatformType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type WeekDay string

const (
	WeekDaySunday    WeekDay = "SUNDAY"
	WeekDayMonday    WeekDay = "MONDAY"
	WeekDayTuesday   WeekDay = "TUESDAY"
	WeekDayWednesday WeekDay = "WEDNESDAY"
	WeekDayThursday  WeekDay = "THURSDAY"
	WeekDayFriday    WeekDay = "FRIDAY"
	WeekDaySaturday  WeekDay = "SATURDAY"
)

var AllWeekDay = []WeekDay{
	WeekDaySunday,
	WeekDayMonday,
	WeekDayTuesday,
	WeekDayWednesday,
	WeekDayThursday,
	WeekDayFriday,
	WeekDaySaturday,
}

func (e WeekDay) IsValid() bool {
	switch e {
	case WeekDaySunday, WeekDayMonday, WeekDayTuesday, WeekDayWednesday, WeekDayThursday, WeekDayFriday, WeekDaySaturday:
		return true
	}
	return false
}

func (e WeekDay) String() string {
	return string(e)
}

func (e *WeekDay) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = WeekDay(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid WeekDay", str)
	}
	return nil
}

func (e WeekDay) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
