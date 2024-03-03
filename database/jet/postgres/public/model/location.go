//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Location struct {
	ID          int64 `sql:"primary_key"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	Type        string
	OwnerID     *int64
	AddressID   int64
	Deleted     *bool
	Status      *LocationStatus
	CreatedBy   *int64
	UpdatedBy   *int64
}
