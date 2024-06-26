//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type LocationStatus string

const (
	LocationStatus_Closed            LocationStatus = "CLOSED"
	LocationStatus_Moved             LocationStatus = "MOVED"
	LocationStatus_Operational       LocationStatus = "OPERATIONAL"
	LocationStatus_TemporarilyClosed LocationStatus = "TEMPORARILY_CLOSED"
)

func (e *LocationStatus) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "CLOSED":
		*e = LocationStatus_Closed
	case "MOVED":
		*e = LocationStatus_Moved
	case "OPERATIONAL":
		*e = LocationStatus_Operational
	case "TEMPORARILY_CLOSED":
		*e = LocationStatus_TemporarilyClosed
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for LocationStatus enum")
	}

	return nil
}

func (e LocationStatus) String() string {
	return string(e)
}
