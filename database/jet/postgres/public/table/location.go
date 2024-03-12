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

var Location = newLocationTable("public", "location", "")

type locationTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz
	Name        postgres.ColumnString
	Description postgres.ColumnString
	Type        postgres.ColumnString
	OwnerID     postgres.ColumnInteger
	AddressID   postgres.ColumnInteger
	Deleted     postgres.ColumnBool
	Status      postgres.ColumnString
	CreatedByID postgres.ColumnInteger
	UpdatedByID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type LocationTable struct {
	locationTable

	EXCLUDED locationTable
}

// AS creates new LocationTable with assigned alias
func (a LocationTable) AS(alias string) *LocationTable {
	return newLocationTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new LocationTable with assigned schema name
func (a LocationTable) FromSchema(schemaName string) *LocationTable {
	return newLocationTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new LocationTable with assigned table prefix
func (a LocationTable) WithPrefix(prefix string) *LocationTable {
	return newLocationTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new LocationTable with assigned table suffix
func (a LocationTable) WithSuffix(suffix string) *LocationTable {
	return newLocationTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newLocationTable(schemaName, tableName, alias string) *LocationTable {
	return &LocationTable{
		locationTable: newLocationTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newLocationTableImpl("", "excluded", ""),
	}
}

func newLocationTableImpl(schemaName, tableName, alias string) locationTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		TypeColumn        = postgres.StringColumn("type")
		OwnerIDColumn     = postgres.IntegerColumn("owner_id")
		AddressIDColumn   = postgres.IntegerColumn("address_id")
		DeletedColumn     = postgres.BoolColumn("deleted")
		StatusColumn      = postgres.StringColumn("status")
		CreatedByIDColumn = postgres.IntegerColumn("created_by_id")
		UpdatedByIDColumn = postgres.IntegerColumn("updated_by_id")
		allColumns        = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, TypeColumn, OwnerIDColumn, AddressIDColumn, DeletedColumn, StatusColumn, CreatedByIDColumn, UpdatedByIDColumn}
		mutableColumns    = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, DescriptionColumn, TypeColumn, OwnerIDColumn, AddressIDColumn, DeletedColumn, StatusColumn, CreatedByIDColumn, UpdatedByIDColumn}
	)

	return locationTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		Type:        TypeColumn,
		OwnerID:     OwnerIDColumn,
		AddressID:   AddressIDColumn,
		Deleted:     DeletedColumn,
		Status:      StatusColumn,
		CreatedByID: CreatedByIDColumn,
		UpdatedByID: UpdatedByIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
