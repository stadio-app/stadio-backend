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

type AuthState struct {
	ID         int64 `sql:"primary_key"`
	LoggedInAt time.Time
	UserID     int64
	IPAddress  *string
}
