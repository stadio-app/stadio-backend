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

type Locations struct {
	ID             int64 `sql:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	Type           string
	OwnerLocations *int64
}
