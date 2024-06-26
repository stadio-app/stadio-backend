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

var AdministrativeDivision = newAdministrativeDivisionTable("public", "administrative_division", "")

type administrativeDivisionTable struct {
	postgres.Table

	// Columns
	CountryCode            postgres.ColumnString
	AdministrativeDivision postgres.ColumnString
	Cities                 postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AdministrativeDivisionTable struct {
	administrativeDivisionTable

	EXCLUDED administrativeDivisionTable
}

// AS creates new AdministrativeDivisionTable with assigned alias
func (a AdministrativeDivisionTable) AS(alias string) *AdministrativeDivisionTable {
	return newAdministrativeDivisionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AdministrativeDivisionTable with assigned schema name
func (a AdministrativeDivisionTable) FromSchema(schemaName string) *AdministrativeDivisionTable {
	return newAdministrativeDivisionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AdministrativeDivisionTable with assigned table prefix
func (a AdministrativeDivisionTable) WithPrefix(prefix string) *AdministrativeDivisionTable {
	return newAdministrativeDivisionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AdministrativeDivisionTable with assigned table suffix
func (a AdministrativeDivisionTable) WithSuffix(suffix string) *AdministrativeDivisionTable {
	return newAdministrativeDivisionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAdministrativeDivisionTable(schemaName, tableName, alias string) *AdministrativeDivisionTable {
	return &AdministrativeDivisionTable{
		administrativeDivisionTable: newAdministrativeDivisionTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newAdministrativeDivisionTableImpl("", "excluded", ""),
	}
}

func newAdministrativeDivisionTableImpl(schemaName, tableName, alias string) administrativeDivisionTable {
	var (
		CountryCodeColumn            = postgres.StringColumn("country_code")
		AdministrativeDivisionColumn = postgres.StringColumn("administrative_division")
		CitiesColumn                 = postgres.StringColumn("cities")
		allColumns                   = postgres.ColumnList{CountryCodeColumn, AdministrativeDivisionColumn, CitiesColumn}
		mutableColumns               = postgres.ColumnList{CountryCodeColumn, AdministrativeDivisionColumn, CitiesColumn}
	)

	return administrativeDivisionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CountryCode:            CountryCodeColumn,
		AdministrativeDivision: AdministrativeDivisionColumn,
		Cities:                 CitiesColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
