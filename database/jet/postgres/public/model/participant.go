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

type Participant struct {
	ID           int64 `sql:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Nickname     *string
	Admin        bool
	Participates bool
	SkillLevel   *string
	EventID      int64
	UserID       int64
}
