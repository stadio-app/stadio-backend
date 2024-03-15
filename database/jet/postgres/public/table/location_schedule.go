//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var LocationSchedule = newLocationScheduleTable("public", "location_schedule", "")

type locationScheduleTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	CreatedAt  postgres.ColumnTimestampz
	UpdatedAt  postgres.ColumnTimestampz
	LocationID postgres.ColumnInteger
	Day        postgres.ColumnString
	On         postgres.ColumnTimestampz
	From       postgres.ColumnTime
	ToDuration postgres.ColumnInteger
	Available  postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LocationScheduleTable struct {
	locationScheduleTable

	EXCLUDED locationScheduleTable
}

// AS creates new LocationScheduleTable with assigned alias
func (a LocationScheduleTable) AS(alias string) *LocationScheduleTable {
	return newLocationScheduleTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LocationScheduleTable with assigned schema name
func (a LocationScheduleTable) FromSchema(schemaName string) *LocationScheduleTable {
	return newLocationScheduleTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LocationScheduleTable with assigned table prefix
func (a LocationScheduleTable) WithPrefix(prefix string) *LocationScheduleTable {
	return newLocationScheduleTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LocationScheduleTable with assigned table suffix
func (a LocationScheduleTable) WithSuffix(suffix string) *LocationScheduleTable {
	return newLocationScheduleTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLocationScheduleTable(schemaName, tableName, alias string) *LocationScheduleTable {
	return &LocationScheduleTable{
		locationScheduleTable: newLocationScheduleTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newLocationScheduleTableImpl("", "excluded", ""),
	}
}

func newLocationScheduleTableImpl(schemaName, tableName, alias string) locationScheduleTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn  = postgres.TimestampzColumn("updated_at")
		LocationIDColumn = postgres.IntegerColumn("location_id")
		DayColumn        = postgres.StringColumn("day")
		OnColumn         = postgres.TimestampzColumn("on")
		FromColumn       = postgres.TimeColumn("from")
		ToDurationColumn = postgres.IntegerColumn("to_duration")
		AvailableColumn  = postgres.BoolColumn("available")
		allColumns       = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, LocationIDColumn, DayColumn, OnColumn, FromColumn, ToDurationColumn, AvailableColumn}
		mutableColumns   = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, LocationIDColumn, DayColumn, OnColumn, FromColumn, ToDurationColumn, AvailableColumn}
	)

	return locationScheduleTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		LocationID: LocationIDColumn,
		Day:        DayColumn,
		On:         OnColumn,
		From:       FromColumn,
		ToDuration: ToDurationColumn,
		Available:  AvailableColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
