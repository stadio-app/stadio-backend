//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var LocationStatus = &struct {
	Closed            postgres.StringExpression
	Moved             postgres.StringExpression
	Operational       postgres.StringExpression
	TemporarilyClosed postgres.StringExpression
}{
	Closed:            postgres.NewEnumValue("CLOSED"),
	Moved:             postgres.NewEnumValue("MOVED"),
	Operational:       postgres.NewEnumValue("OPERATIONAL"),
	TemporarilyClosed: postgres.NewEnumValue("TEMPORARILY_CLOSED"),
}
