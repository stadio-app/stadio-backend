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

type LocationSchedule struct {
	ID         int64 `sql:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LocationID int64
	Day        WeekDay
	On         *time.Time
	From       *time.Time
	ToDuration *int32
	Available  bool
}
