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

type Owners struct {
	ID         int64 `sql:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FirstName  string
	MiddleName *string
	LastName   string
	FullName   string
	IDURL      string
	Verified   bool
	UserOwner  int64
}
