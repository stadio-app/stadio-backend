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

var Country = newCountryTable("public", "country", "")

type countryTable struct {
	postgres.Table

	// Columns
	Code postgres.ColumnString
	Name postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CountryTable struct {
	countryTable

	EXCLUDED countryTable
}

// AS creates new CountryTable with assigned alias
func (a CountryTable) AS(alias string) *CountryTable {
	return newCountryTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CountryTable with assigned schema name
func (a CountryTable) FromSchema(schemaName string) *CountryTable {
	return newCountryTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CountryTable with assigned table prefix
func (a CountryTable) WithPrefix(prefix string) *CountryTable {
	return newCountryTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CountryTable with assigned table suffix
func (a CountryTable) WithSuffix(suffix string) *CountryTable {
	return newCountryTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCountryTable(schemaName, tableName, alias string) *CountryTable {
	return &CountryTable{
		countryTable: newCountryTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newCountryTableImpl("", "excluded", ""),
	}
}

func newCountryTableImpl(schemaName, tableName, alias string) countryTable {
	var (
		CodeColumn     = postgres.StringColumn("code")
		NameColumn     = postgres.StringColumn("name")
		allColumns     = postgres.ColumnList{CodeColumn, NameColumn}
		mutableColumns = postgres.ColumnList{NameColumn}
	)

	return countryTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Code: CodeColumn,
		Name: NameColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
